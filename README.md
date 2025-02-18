# Weather Forecast Application
A simple weather forecast application that fetches weather data from the AccuWeather API and displays it in a web interface. The application consists of a Go Fiber backend and a Node.js frontend, deployed on Kubernetes.


# Features
- Backend: Fetches weather data from the AccuWeather API using a dynamic locationKey.

- Frontend: Allows users to input a locationKey and view the weather forecast.

- Deployment: Containerized using Docker and deployed on Kubernetes.

# Technologies Used
- Backend: Go (Golang) with Fiber

- Frontend: Node.js with Express

- Containerization: Docker

- Orchestration: Kubernetes (K8s)

- API: AccuWeather API

# Prerequisites
Before running the application, ensure you have the following installed:

1. Go (for backend development)

2. Node.js (for frontend development)

3. Docker (for containerization)

4. Kubernetes (for deployment)

5. kubectl (Kubernetes CLI)

6. Minikube (for local Kubernetes cluster, optional)

7. AccuWeather API Key (sign up at AccuWeather Developer Portal)

# Setup Instructions
1. Clone the Repository
```shell
git clone https://github.com/your-username/weather-forecast-app.git
cd weather-forecast-app
```
2. Configure the Backend
- Open the `backend/main.go` file.

- Replace `YOUR_ACCUWEATHER_API_KEY` with your actual AccuWeather API key.

3. Build Docker Images
#### Backend
1. Navigate to the backend directory:

```shell
cd backend
```
2. Build the Docker image:

```shell
docker build -t your-dockerhub-username/backend:latest -f Dockerfile .
```
3. Push the image to Docker Hub:

```shell
docker push your-dockerhub-username/backend:latest
```

#### Frontend
1. Navigate to the frontend directory:

```shell
cd ../frontend
```
2. Build the Docker image:

```shell
docker build -t your-dockerhub-username/frontend:latest -f Dockerfile .
```
3. Push the image to Docker Hub:

```shell
docker push your-dockerhub-username/frontend:latest
```
4. Deploy to Kubernetes
- Apply the backend deployment:
```shell
kubectl apply -f backend-deployment.yaml
```
- Apply the frontend deployment:
```shell
kubectl apply -f frontend-deployment.yaml
```
- Verify the deployments
```shell
kubectl get pods
kubectl get svc
```
5. Access the Application
#### Using Minikube
1. Start Minikube tunnel:
```shell
minikube tunnel
```
2. Get the external IP of `frontend-service`
```shell
kubectl get svc frontend-service
```
3. Open the external IP in browser.
#### Using NodePort
1. Get the NodePort assigned to the frontend-service:
```shell
kubectl get svc frontend-service
```
2. Access the frontend using:
```shell
http://<node-ip>:<node-port>
```
- `<node-ip>`: The IP of your Kubernetes node (e.g., localhost for Minikube).

- `<node-port>`: The port assigned by Kubernetes.

# Usage
1. Open the frontend in your browser.

2. Enter a valid locationKey (e.g., 12345 for New York) and click Get Weather.

3. The weather data for the specified location will be displayed.

# Project Structure
```shell
weather-forecast-app/
├── backend/
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   └── Dockerfile
├── frontend/
│   ├── server.js
│   ├── package.json
│   └── Dockerfile
├── backend-deployment.yaml
├── frontend-deployment.yaml
└── README.md
```
# Troubleshooting
1. Failed to Fetch Weather Data
- Ensure the backend is running and accessible.
- Check the backend logs for errors:
```shell
kubectl logs <backend-pod-name>
```
- Verify the `apiKey` and `locationKey` are correct.
2. Frontend Not Accessible
- Ensure the frontend-service is running.
- Check the frontend logs for errors:
```shell
kubectl logs <frontend-pod-name>
```
- Use kubectl port-forward to access the frontend locally:
```shell
kubectl port-forward svc/frontend-service 4000:80
```
# Contributing
Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (git checkout -b feature/your-feature).
3. Commit your changes (git commit -m 'Add some feature').
4. Push to the branch (git push origin feature/your-feature).
5. Open a pull request.
# License
This project is licensed under the MIT License. See the LICENSE file for details.

# Acknowledgments
- [AccuWeather](https://www.accuweather.com/) for providing the weather data API.
- [Go Fiber](https://gofiber.io/) for the lightweight web framework.
- [Express](https://expressjs.com/) for the Node.js web framework.