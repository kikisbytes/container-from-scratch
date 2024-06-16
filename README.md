## Getting Started

### Prerequisites
1. Set up VM on your local machine (mac) - https://www.youtube.com/watch?v=O19mv1pe76M
   - This video will also show you know to share files between the VM and your host machine. This makes it easier to run our go code.
2. Install Go - https://golang.org/doc/install

### Setting Up the project
1. Clone this project to the dir that you shared with your VM. 
2. Start up the VM
3. Create a copy of the ubuntu filesystem - https://askubuntu.com/questions/1049930/how-to-copy-root-file-system-in-ubuntu
```
rsync -aAXv / --exclude={"/dev/*","/proc/*","/sys/*","/tmp/*","/run/*","/mnt/*","/media/*","/lost+found","/home/*"} /docker-fs
```
- This will create a copy of the ubuntu filesystem in the /docker-fs directory. This will allow us to run the project in a container that has the same filesystem as our VM.

Then we want to create a basic express app inside the /docker-fs directory.
- cd /
- cd docker-fs
- mkdir express-app
- cd express-app
- create the following 3 files:

index.js
```js
const express = require('express');
const app = express();
const port = 3000;

app.get('/', (req, res) => {
  res.send('Hello, World! \n');
});

app.listen(port, () => {
  console.log(`App running at http://localhost:${port}`);
});
```

package.json
```json
{
  "name": "express-app",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "start": "node index.js"
  },
  "dependencies": {
    "express": "^4.19.2"
  },
  "author": "Kiki's Bytes",
  "license": "ISC"
}
```

Dockerfile
```Dockerfile
FROM node:v18
WORKDIR /express-app
RUN npm install
CMD ["node", "index.js"]
```

## To start our container
1. cd into the project directory. 
2. Start the project by running the following command. This will start a container that will install nodejs and run our express app.
```
sudo go run main.go run /bin/bash
```
3. Open a new terminal on the VM and test our express app by running the following command. 
```
curl localhost:3000
```
