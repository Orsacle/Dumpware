package installer

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	phpVersion    = "8.4.22"
	phpVCVersion  = "vs17"
	phpArch       = "x64"
	phpInstallDir = `C:\tools\php`
)

func phpDownloadCandidates() []string {
	filename := fmt.Sprintf("php-%s-Win32-%s-%s.zip", phpVersion, phpVCVersion, phpArch)
	return []string{
		"https://downloads.php.net/~windows/releases/" + filename,
		"https://downloads.php.net/~windows/releases/archives/" + filename,
	}
}

func InstallPHPWindows() error {
	var lastErr error
	for _, url := range phpDownloadCandidates() {
		fmt.Printf("Trying %s ...\n", url)
		if err := downloadAndExtractZip(url, phpInstallDir); err != nil {
			lastErr = err
			continue
		}
		fmt.Printf("PHP %s installed to %s\n", phpVersion, phpInstallDir)
		if err := addToUserPath(phpInstallDir); err != nil {

			fmt.Fprintf(os.Stderr, "PHP installed, but could not update PATH automatically: %v\n", err)
			fmt.Printf("Add %s to your PATH manually to use the php command.\n", phpInstallDir)
		}
		return nil
	}
	return fmt.Errorf("could not download PHP from any known location: %w", lastErr)
}

func downloadAndExtractZip(url, destDir string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed with status %s", resp.Status)
	}

	tmpFile, err := os.CreateTemp("", "php-*.zip")
	if err != nil {
		return fmt.Errorf("could not create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		return fmt.Errorf("could not save download: %w", err)
	}

	return unzip(tmpFile.Name(), destDir)
}

func unzip(zipPath, destDir string) error {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return fmt.Errorf("could not open zip: %w", err)
	}
	defer r.Close()

	if err := os.MkdirAll(destDir, 0o755); err != nil {
		return fmt.Errorf("could not create install dir: %w", err)
	}

	for _, f := range r.File {
		fpath := filepath.Join(destDir, f.Name)

		rel, err := filepath.Rel(destDir, fpath)
		if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path in zip: %s", f.Name)
		}

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(fpath, 0o755); err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), 0o755); err != nil {
			return err
		}

		if err := extractFile(f, fpath); err != nil {
			return err
		}
	}

	return nil
}

func extractFile(f *zip.File, destPath string) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	outFile, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, rc)
	return err
}

func addToUserPath(dir string) error {
	currentUserPath, err := userPath()
	if err != nil {
		return fmt.Errorf("could not read user PATH: %w", err)
	}

	for _, p := range filepath.SplitList(currentUserPath) {
		if strings.EqualFold(p, dir) {
			return nil
		}
	}

	newPath := dir
	if currentUserPath != "" {
		newPath = currentUserPath + ";" + dir
	}

	cmd := exec.Command("powershell", "-NoProfile", "-Command",
		"[Environment]::SetEnvironmentVariable('PATH', $env:DUMPWARE_NEW_PATH, 'User')")
	cmd.Env = append(os.Environ(), "DUMPWARE_NEW_PATH="+newPath)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("could not set user PATH: %w (%s)", err, string(out))
	}
	return nil
}

func userPath() (string, error) {
	cmd := exec.Command("powershell", "-NoProfile", "-Command",
		"[Environment]::GetEnvironmentVariable('PATH', 'User')")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

// php is really fucking special holy
