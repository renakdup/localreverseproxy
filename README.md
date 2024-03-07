# Local Reverse Proxy
The reverse proxy server for the local environment that can help to run several docker projects locally via the 80 port.

## How it works (diagram)
```bash
      Client              Host (OS mapping)        Reverse Proxy        Target Server (Docker App)
     +------+              +-------------+       +--------------+         +----------------+
     |      | example.loc  |             |       │              │  proxy  │                │
     |      |─────────────►| example.loc │──────►│              │────────►│                │
     |      |              |             │       │ 127.0.0.1:80 │         │ localhost:8080 │
     |      |◄─────────────| 127.0.0.1   │◄──────│              │◄────────│                │
     |      |              |             │       │              │         │                │
     +------+              +-------------+       +--------------+         +----------------+
     (Browser)               /etc/hosts
```

## How to use it
1. [Download and Install Go](https://go.dev/dl/)
2. Clone this repo where you want.
3. Configure ports for your docker projects:
   - Create `docker-compose.override.yml` in a project.
   - Add there apache or nginx service and change a port. Because by default most of projects uses 80 or 443 ports.  
     Yaml should look like this:
   ```yaml
   services:   
     nginx:
       ports: !override # helps to reset ports. Works only on the latest version of docker-compose
         - "8080:80" # Now set up new external port `8080`
   ```
   - Recreate docker containers into your project directory.
   ```bash
   docker-compose up -d --force-recreate
   ```
   - Do the same steps for another docker project but use another port, as example `8090`.   
   **Remember one port for one service/device!**
> [!NOTE]
> Your 2 and more projects have to run and work without any issues, if not check that external ports be unique.

4. Go to the directory of cloned `localreverseproxy` repository.
5. Add your docker services into `services.json` file. Where:
   - `Key` is local domain address.
   - `Value` is URL:Port of docker service. Use `http://localhost:<your_service_port>`
   
> [!NOTE]
> You can use simple regexp in `Key` fields of `services.json` file.  
> To have a mask, like: `*.local`.

> [!NOTE]
> If you change the `services.json` file you don't need to compile app again. It is read in runtime.

> [!NOTE]
> Don't forget add domains into your hosts file.

6. Let's compile Go app
```bash
go mod tidy
go build
```
7. Run Local Reverse Proxy
```bash
./localreverseproxy
```
If everything ok you will see in a console
```bash
Reverse proxy is started.
```

---

## Roadmap
- [x] Configure routing via `json`.
- [x] Implement reverse-proxy functionality.
- [x] Add documentation how to use it.
- [x] Show error if port is busy.
- [x] Add requests logging in console.
- [ ] Add service response logging in console. 
- [ ] Add supporting url paths.
- [ ] Add supporting https.
