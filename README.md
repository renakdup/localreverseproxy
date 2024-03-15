# Local Reverse Proxy
The reverse proxy server for the local environment can help to run several docker projects locally via the 80 port.

Each time when you face issues when you need to run several docker projects and you don't want to stop and run it again and again, or you just need your projects to work on 80 port. Use this Local Reverse Proxy!

Article: [How to run multiple Docker projects under port 80 using an off-the-shelf local Reverse Proxy](https://wp-yoda.com/en/environment/how-to-run-multiple-docker-projects-under-port-80-using-an-off-the-shelf-local-reverse-proxy/).

## How it works (diagram)
```text
    Client                  Host (OS mapping)          Reverse Proxy        Target Server (Docker App)
   +------+    request      +-----------------+       +--------------+         +----------------+
   │      │ 'example.local' │                 │       │              │  proxy  │                │
   │      │────────────────►│ 'example.local' │──────►│              │────────►│                │
   │      │    response     │    mapped to    │       │ 127.0.0.1:80 │         │ localhost:8080 │
   │      │◄────────────────│   '127.0.0.1'   │◄──────│              │◄────────│                │
   │      │                 │                 │       │              │         │                │
   +------+                 +-----------------+       +--------------+         +----------------+
   (Browser)                    /etc/hosts
```

## How to set up Local Reverse Proxy

### How run Local Reverse Proxy
1. [Download and Install Go](https://go.dev/dl/)  
   - For macOS/Linux just follow instruction. 
   - For windows - install Go into Windows OS, and use `CMD` to build and run Go app.
2. Clone and go into cloned repository.
3. Let's compile Go app 
   ```bash
   go mod tidy
   go build
   ```
4. Run Local Reverse Proxy:
   ```bash
   ./localreverseproxy

   # or if you have permissions errors use with sudo
   sudo ./localreverseproxy 
   ```
5. If everything ok you will see in a console:
   ```bash
   Reverse proxy is started.
   ```
6. Now visit `http://localhost`, you should see in the browser
   ```bash
   Server proxy started successfully!
   ```
   and in the console you should see
   ```bash
   Requested `localhost` host with method `GET` and URI `/`
   Requested `localhost` host with method `GET` and URI `/favicon.ico`
   ```
   
   > [!NOTE]
   > Each http request will be logged and you will see it in the console, so you can check how things going.

### How to configure docker projects to work with Local Reverse Proxy
1. Configure ports for your docker projects:
   - Create `docker-compose.override.yml` in your projects.
   - Add there apache or nginx service and change a port. Because by default most of projects uses 80 or 443 ports.  
     Yaml should look like this:
   ```yaml
   services:   
     nginx:
       ports: !override # helps to reset ports.
         - "8080:80" # Now set up new external port `8080`
   ```
   > [!NOTE]
   > `!override` works only on the version v2.24.5 and above of docker-compose.

   For another project you can set 8090 port.
   - Recreate docker containers inside your project directory to containers use new port.
   ```bash
   docker-compose up -d --force-recreate
   ```
   - Do the same steps for another docker project but use another port, as example `8090`.   
   
   > [!NOTE]
   > **Remember one port for one service/container/device!**

   > [!NOTE]
   > Your 2 and more projects have to run and work without any issues, if not check that external ports be unique.
   
2. Go to the directory of cloned `localreverseproxy` repository again.
3. Add your docker projects domains into `services.json` file. Where:
   - `Key` is local domain address.
   - `Value` is URL:Port of docker service. Use `http://localhost:<your_service_port>`
   
> [!NOTE]
> You can use simple regexp in `Key` fields of `services.json` file.  
> To have a mask, like: `*.local`.

> [!NOTE]
> **If you change the `services.json` file you don't need to compile app again. It is read in runtime.**

> [!NOTE]
> Don't forget add domains into your `hosts` file.

---

## Roadmap
- [x] Configure routing via `json`.
- [x] Implement reverse-proxy functionality.
- [x] Add documentation how to use it.
- [x] Show error if port is busy.
- [x] Add requests logging in console.
- [x] Add service response logging in console. 
- [ ] Add supporting https
- [ ] Add supporting url paths
