import express from 'express';
import cors from 'cors'

const app = express();

app.use(express.json(), cors({
    methods: ['GET', 'POST', 'PUT', 'DELETE'],

}));


export default app;




