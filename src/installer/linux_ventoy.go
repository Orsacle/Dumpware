// no idea how i could've done this. this is all ai-generated!!

package installer

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	ventoyLatestReleaseAPI = "https://api.github.com/repos/ventoy/Ventoy/releases/latest"
	ventoyInstallDir       = "/opt/ventoy"
)

type githubRelease struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

// InstallVentoyLinux downloads the latest Ventoy Linux release tarball from
// GitHub and extracts it to /opt/ventoy. Used regardless of package manager
// since no distro packages this natively (Arch's AUR build is the closest
// thing, but Dumpware doesn't shell out to an AUR helper).
func InstallVentoyLinux() error {
	release, err := fetchLatestVentoyRelease()
	if err != nil {
		return fmt.Errorf("could not fetch latest Ventoy release info: %w", err)
	}

	var assetURL string
	for _, a := range release.Assets {
		if strings.HasSuffix(a.Name, "-linux.tar.gz") {
			assetURL = a.BrowserDownloadURL
			break
		}
	}
	if assetURL == "" {
		return fmt.Errorf("no linux.tar.gz asset found in release %s", release.TagName)
	}

	tmpFile, err := os.CreateTemp("", "ventoy-*.tar.gz")
	if err != nil {
		return fmt.Errorf("could not create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	if err := downloadFile(assetURL, tmpFile.Name()); err != nil {
		return fmt.Errorf("could not download %s: %w", assetURL, err)
	}

	if err := extractTarGz(tmpFile.Name(), ventoyInstallDir); err != nil {
		return fmt.Errorf("could not extract Ventoy archive: %w", err)
	}

	version := strings.TrimPrefix(release.TagName, "v")
	script := filepath.Join(ventoyInstallDir, "ventoy-"+version, "Ventoy2Disk.sh")
	_ = os.Chmod(script, 0o755) // best effort, tar usually preserves the exec bit already

	fmt.Printf("Ventoy %s extracted to %s\n", release.TagName, ventoyInstallDir)
	fmt.Printf("To write it to a USB drive, run (as root): %s -i /dev/sdX\n", script)
	return nil
}

func fetchLatestVentoyRelease() (*githubRelease, error) {
	resp, err := http.Get(ventoyLatestReleaseAPI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status %s", resp.Status)
	}

	var release githubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, fmt.Errorf("could not parse GitHub API response: %w", err)
	}
	return &release, nil
}

func extractTarGz(archivePath, destDir string) error {
	f, err := os.Open(archivePath)
	if err != nil {
		return err
	}
	defer f.Close()

	gzr, err := gzip.NewReader(f)
	if err != nil {
		return fmt.Errorf("could not open gzip stream: %w", err)
	}
	defer gzr.Close()

	if err := os.MkdirAll(destDir, 0o755); err != nil {
		return fmt.Errorf("could not create install dir: %w", err)
	}

	tr := tar.NewReader(gzr)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		fpath := filepath.Join(destDir, hdr.Name)

		rel, err := filepath.Rel(destDir, fpath)
		if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path in archive: %s", hdr.Name)
		}

		switch hdr.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(fpath, 0o755); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(fpath), 0o755); err != nil {
				return err
			}
			if err := extractTarFile(tr, fpath, hdr.Mode); err != nil {
				return err
			}
		}
	}
	return nil
}

func extractTarFile(tr *tar.Reader, destPath string, mode int64) error {
	out, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(mode))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, tr)
	return err
}
