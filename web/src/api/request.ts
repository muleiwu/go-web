// src/axios.js
import axios from 'axios';

// console.log(import.meta.env)
const instance = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL, // 替换为你的API基础URL
    timeout: 1000,
    // headers: {'X-Custom-Header': 'foobar'}
});


export default instance;
