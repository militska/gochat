FROM redis:latest

RUN echo "vm.overcommit_memory = 1" >> /etc/sysctl.conf
RUN sysctl vm.overcommit_memory=1

EXPOSE 511
EXPOSE 6379