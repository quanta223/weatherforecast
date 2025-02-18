const express = require('express');
const axios = require('axios');

const app = express();
const port = 4000;

// Serve a simple HTML form
app.get('/', (req, res) => {
    res.send(`
        <h1>Weather Forecast</h1>
        <form action="/weather" method="get">
            <label for="locationKey">Enter Location Key:</label>
            <input type="text" id="locationKey" name="locationKey" required>
            <button type="submit">Get Weather</button>
        </form>
    `);
});

// Fetch weather data from the backend
app.get('/weather', async (req, res) => {
    const locationKey = req.query.locationKey;
    if (!locationKey) {
        return res.status(400).send('Location key is required');
    }

    try {
        const response = await axios.get(`http://backend-service:3000/weather?locationKey=${locationKey}`);
        res.send(`<h1>Weather Forecast</h1><pre>${JSON.stringify(response.data, null, 2)}</pre>`);
    } catch (error) {
        res.status(500).send('Failed to fetch weather data');
    }
});

// Start the server
app.listen(port, () => {
    console.log(`Frontend listening at http://localhost:${port}`);
});