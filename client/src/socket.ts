import { io } from 'socket.io-client';

// "undefined" means the URL will be computed from the `window.location` object
const URL: string | undefined = process.env.NODE_ENV === 'production' ? undefined : 'http://localhost:8080';

// @ts-ignore
export const socket = io(URL, {
    autoConnect: false,
    cors: {
        origin: "https://localhost:3000",
        methods: ["GET", "POST"],
        allowedHeaders: ["Access-Control-Allow-Origin"],
        credentials: true
      }
});