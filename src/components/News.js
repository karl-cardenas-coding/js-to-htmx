/**
 * Copyright (c) karl-cardenas-coding
 * SPDX-License-Identifier: MIT
 */

import React, { useState, useEffect } from 'react';
import './News.css';

async function getNews() {
    const url = "https://min-api.cryptocompare.com/data/v2/news/?lang=EN";
    
    try {
        const response = await fetch(url);
        const data = await response.json();
        // Return only the top five news articles in the data list
        if (data.Data.length > 5) {
            return data.Data.slice(0, 5);
        }
        return data.Data;
    }
    catch (error) {
        console.error(error);
        return [];
    }
}

function News() {
    const [news, setNews] = useState([]);

    useEffect(() => {
        async function fetchNews() {
            const fetchedNews = await getNews();
            setNews(fetchedNews);
        }

        fetchNews();
    }, []);

    return (
        <div class="news-section">
            <h2>Latest News</h2>
            <div id="news-list">
                <ul>
                    {news.map((article, index) => (
                        <li key={index}>
                            <a href={article.url} target="_blank" rel="noopener noreferrer">
                                <strong>{article.title}</strong> - {new Date(article.published_on * 1000).toLocaleDateString()}
                            </a>
                        </li>
                    ))}
                </ul>
            </div>
        </div>
    );
}

export default News;
