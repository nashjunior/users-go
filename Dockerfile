FROM golang:latest

RUN apt update && apt install -y --no-install-recommends \
    git \
    ca-certificates  ca-certificates-java \
    default-jre \
    zsh \
    curl \
    wget

RUN useradd -m -u 1000 go

USER go
ENV JAVA_HOME="/usr/lib/jvm/java-17-openjdk-amd64/"

WORKDIR /home/go/app

RUN bash -c "$(curl --fail --show-error --silent --location https://raw.githubusercontent.com/zdharma-continuum/zinit/HEAD/scripts/install.sh)" -- \
    bash -c zinit self-update

RUN echo "zinit light zdharma/fast-syntax-highlighting" >>~/.zshrc && \
    echo "zinit light zsh-users/zsh-autosuggestions" >>~/.zshrc && \
    echo "zinit light zsh-users/zsh-completions" >>~/.zshrc && \
    echo "zinit ice depth=1; zinit light romkatv/powerlevel10k" >>~/.zshrc && \
    echo 'HISTFILE=~/zsh/.zsh_history\nHISTSIZE=10000' >>~/.zshrc 


CMD ["sh", "-c", "tail -f /dev/null" ]

