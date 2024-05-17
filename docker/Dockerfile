from golang:latest
WORKDIR /app
RUN mkdir -p /app/node
RUN git clone https://github.com/swanchain/go-computing-provider.git && cd go-computing-provider
RUN cd go-computing-provider && git checkout releases
RUN cd go-computing-provider &&  make
ENV CP_PATH=/app/node
RUN sed -i 's/sudo //g' /app/go-computing-provider/install.sh
RUN apt update
RUN apt install python3-full python3-pip -y 
ENV PIP_BREAK_SYSTEM_PACKAGES=1
RUN cd go-computing-provider &&  bash install.sh
ENV PATH=$PATH:/app/go-computing-provider
RUN apt-get clean
EXPOSE 8085
