# eb-fasthttp-sample
Sample FastHTTP application for AWS Elastic Beanstalk

build.sh should executable

`chmod +x filename.sh`

# Install Elastic Beanstalk command line interface

`curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"`

`sudo apt-get install python3.7`

`curl -O https://bootstrap.pypa.io/get-pip.py`

`python3.7 get-pip.py --user`

`pip install awsebcli --upgrade --user`

export PATH=~/.local/bin/:$PATH

source ~/.profile

# Initialize and deploy

`eb init`

`eb create`

`eb deploy`


<a title="Deploy to AWS" href="https://console.aws.amazon.com/elasticbeanstalk/home?region=eu-central-1#/newApplication?applicationName=FastHTTPServer&solutionStackName=Go&tierName=WebServer" target="_blank"><img src="http://d0.awsstatic.com/product-marketing/Elastic%20Beanstalk/deploy-to-aws.png" height="40"></a>
