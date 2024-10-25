FROM golang:1.23.2

# Install golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
    sh -s -- -b $(go env GOPATH)/bin v1.61.0

# Install dependencies for GLFW
RUN apt-get update && \
	apt-get install -y libx11-dev \
	    libgl1-mesa-dev \
        libasound2-dev \
        libgles2-mesa-dev \
        libalut-dev \
        libxcursor-dev \
        libxi-dev \
        libxinerama-dev \
        libxrandr-dev \
        libxxf86vm-dev