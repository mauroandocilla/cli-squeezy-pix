const os = require('os');
const https = require('https');
const fs = require('fs');
const path = require('path');

function download(url, destPath, callback) {
  https.get(url, (response) => {
    // If a 302 status code is received (redirection), follow the new location
    if (response.statusCode === 302 && response.headers.location) {
      console.log(`Redirecting to ${response.headers.location}`);
      download(response.headers.location, destPath, callback); // Recursively follow the redirection
    } else if (response.statusCode === 200) {
      // If a 200 status code is received, download the file
      const file = fs.createWriteStream(destPath);
      response.pipe(file);
      file.on('finish', () => {
        file.close(callback);
      });
    } else {
      console.error(`Failed to download binary: ${response.statusCode}`);
      process.exit(1);
    }
  }).on('error', (err) => {
    console.error(`Error downloading binary: ${err.message}`);
    process.exit(1);
  });
}

// Read the base release URL from the file
const releaseUrl = fs.readFileSync(path.join(__dirname, 'release-url.env'), 'utf8').split('=')[1].trim();

// Detect the platform and build the full binary URL
const platform = os.platform();
const arch = os.arch();

console.log(`Detected platform: ${platform}, architecture: ${arch}`);

let binaryName = '';

if (platform === 'darwin' && (arch === 'x64' || arch === 'arm64')) {
  binaryName = 'spix-macos';
} else if (platform === 'linux' && arch === 'x64') {
  binaryName = 'spix-linux';
} else if (platform === 'win32' && arch === 'x64') {
  binaryName = 'spix.exe';
} else {
  console.error(`Unsupported platform or architecture: ${platform} ${arch}`);
  process.exit(1);
}

const url = `${releaseUrl}/${binaryName}`;
const destPath = path.join(__dirname, 'spix');

// Download the binary from GitHub Releases
download(url, destPath, () => {
  console.log(`Downloaded and installed spix binary at ${destPath}`);

  // If Unix, make the binary executable
  if (platform !== 'win32') {
    fs.chmodSync(destPath, '755');
  }

  // If Windows, make sure the file has the .exe extension
  if (platform === 'win32') {
    fs.renameSync(destPath, destPath + '.exe');
  }
});
