with import <nixpkgs> {};
mkShell {
  nativeBuildInputs = [
    pkgs.go
    pkgs.cobra-cli
  ];
}
