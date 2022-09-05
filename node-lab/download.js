var http = require("http");
var https = require('https');

class HttpsWithProxyClient {
    constructor(proxyConf) {
        this.proxyConf = proxyConf;
        this.proxyAgent = null;
        this.lastRequestHostname = null;
    }

    // Async private method, used only internally by HttpsWithProxyClient
    // Handle the connection with the HTTP proxy
    // In case the connection is successful, it returns a new http.Agent with keepAlive activated
    async _connectToProxy(url) {
        return new Promise((resolve, reject) => {
            const headers = {
                // 'Proxy-Authorization': 'Basic ' + Buffer.from(this.proxyConf.proxy_username + ':' + this.proxyConf.proxy_password).toString('base64')
            }

            const urlParsed = new URL(url);
            http.request({ // establishing a tunnel
                host: this.proxyConf.proxy_host,
                port: this.proxyConf.proxy_port,
                method: 'CONNECT',
                path: `${urlParsed.hostname}:443`,
                headers
            }).on('connect', (res, socket) => {
                if (res.statusCode === 200) {
                    resolve(new https.Agent({ socket: socket, keepAlive: true }));
                } else {
                    reject('Could not connect to proxy!')
                }

            }).on('error', (err) => {
                reject(err.message);
            }).on('timeout', (err) => {
                reject(err.message);
            }).end();
        });
    }

    // Public asynchronous method used to make an HTTPs get request on a given URL through an HTTP proxy
    async getURL(url, headers) {
        return new Promise(async (resolve, reject) => {
            const urlParsed = new URL(url);
            headers = headers || {};

            // If there's no current valid connection established with the proxy
            // or if the host linked to the URL requested is different from the previous request
            // -> recreate a connection with the proxy
            if (!this.proxyAgent || this.lastRequestHostname !== urlParsed.hostname) {
                try {
                    this.proxyAgent = await this._connectToProxy(url);
                    this.lastRequestHostname = urlParsed.hostname;
                } catch (e) {
                    return reject(e);
                }

            }

            const req = https.get({
                host: urlParsed.hostname,
                path: urlParsed.pathname,
                agent: this.proxyAgent,
                headers: headers,
            }, resolve);

            req.on('error', (err) => {
                this.proxyAgent = null;
                reject(err.message);
            })
        })
    }
}

const httpsClient = new HttpsWithProxyClient({
    // "proxy_username": 'xxxxx',
    // "proxy_password": 'xxxxxx',
    "proxy_host": '127.0.0.1',
    "proxy_port": 7890
});

(async () => {


    try {


        // Side effect is even if you have a rotating proxy, the IP will be stable
        
        const url = "https://github.com/jitsi/webrtc/releases/download/v94.0.0/WebRTC.xcframework.tgz";


        const res = await httpsClient.getURL(url);
        console.log(res.statusCode);
        console.log(res.headers.location);
        // console.log(res.body);
    } catch (e) {
        console.error(e)
    }

})();