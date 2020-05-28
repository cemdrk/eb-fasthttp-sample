# eb-fasthttp-sample
Sample FastHTTP application for AWS Elastic Beanstalk


chmod +x build.sh

curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"

sudo apt-get install python3.7

curl -O https://bootstrap.pypa.io/get-pip.py

python3.7 get-pip.py --user

pip install awsebcli --upgrade --user

export PATH=~/.local/bin/:$PATH

source ~/.profile

eb init

eb create

eb deploy
