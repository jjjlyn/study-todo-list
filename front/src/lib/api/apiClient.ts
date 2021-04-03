import axios from "axios";

const { location } = window;

const apiClient = axios.create({
    // baseURL: "http://localhost:3001/api",
    baseURL: `${location.protocol}//${location.host}/api`,
});

export default apiClient;