#!/bin/bash

git config core.hooksPath .githooks

if ! command -v golangci-lint &> /dev/null
then
    echo "golangci-lint não encontrado. Instalando..."
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest
    export PATH="$HOME/.golangci-lint/bin:$PATH"
else
    echo "golangci-lint já está instalado."
fi

echo "Hooks configurados para usar .githooks"
