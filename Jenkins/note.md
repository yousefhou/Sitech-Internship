# CI/CD :
its a way to take code package it up and deploy it to a system , that system could be a serverless like a lambda function or azure function or virtual machine or an EC2 instance , or container running somewhare with  docker .
# CI : 
continuous integration its you take the code and package it up like a gift you wrapping (the gifts it comes pieces you package it up and you wrapping it with gifts paper ) you taking the code and you packaging it up and then you giving it to CD process .

# CD : 
continuous deployment/delivery is where you deploy the code to some system (cloud be a serverless like a lambda ...) you take that gift you wrapped you put it in your car you drive to that person house and you deliver it to their house , once the cod eis packeged and tested and all is good in the CI process you then deliver it to the system you're running on.

# CD(delivery) vs CD(deployment)  :
the key differences here delivery there is some button that you have to press to deploy the code .
and continuous deployment its all about just everything is automatic (zero human intervention ).

 # why Jenkins :
free and a lot of plugins and enterprise are available too .

#  how to insatll Jenkins :
> wget -q -O - https://pkg.jenkins.io/debian-stable/jenkins.io.key |sudo gpg --dearmor -o /usr/share/keyrings/jenkins.gpg
> sudo sh -c 'echo deb [signed-by=/usr/share/keyrings/jenkins.gpg] http://pkg.jenkins.io/debian-stable binary/ > /etc/apt/sources.list.d/jenkins.list'
>  sudo apt update
>  sudo apt install jenkins

# start Jenkins :
> sudo systemctl start jenkins.service
> sudo systemctl status jenkins

# open firewall :
> sudo ufw allow 8080

# Check ufw’s status to confirm the new rules: 
>  sudo ufw status

# to know the passwrod od Jenkins:
> sudo cat /var/lib/jenkins/secrets/initialAdminPassword


#  plugins allow you to connect to another services.

#  how to restart jenkins (if need to for a new plugin):
> sudo systemctl restart jenkins
 


