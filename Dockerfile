FROM jenkins/jenkins:latest
RUN /usr/local/bin/install-plugins.sh cloudbees-folder ssh-slaves

