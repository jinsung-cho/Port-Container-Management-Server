import axios from "axios";
const axiosInstance = axios.create({
    baseURL: "http://10.0.2.180:4000",
    headers: {
        "Content-Type": "application/json",
        'Authorization': `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidGVzdCJ9.Ud1nm095AmrhaYFrgDPHvBlm3W7GQ7fz7xUQbG8aZ-U`,
    },
});

export default axiosInstance;