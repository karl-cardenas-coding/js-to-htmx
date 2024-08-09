import './App.css';
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import BitcoinPrice from './components/BitcoinPrice';
import Ticker from './components/Ticker';
import bitcoinLogo from './bitcoin.png';
import ethcoinLogo from './ethereum.png';
import usdccoinLogo from './usdc.png';

function App() {
  const [portfolio, setPortfolio] = useState({});
  const [selectedCoin, setSelectedCoin] = useState({
    name: 'Bitcoin',
    title: 'Bitcoin',
    price: 'Loading...',
    lastUpdate: '12:34:5678',
    icon: bitcoinLogo,
  });

  async function fetchExchangeRates() {
    const assets = ['BTC', 'ETH', 'USDC'];
    const promises = assets.map(asset => {
      const fsym = asset === 'USDC' ? 'USDT' : asset; // Use USDT for USDC
      return axios.get(`https://min-api.cryptocompare.com/data/price?fsym=${fsym}&tsyms=USD`);
    });
    const responses = await Promise.all(promises);
    const exchangeRates = responses.reduce((acc, response, index) => {
      const key = assets[index];
      acc[key] = response.data.USD;
      return acc;
    }, {});
    setPortfolio(exchangeRates);
  }

  useEffect(() => {
    // Initial fetch
    fetchExchangeRates();

    // Set interval for fetching data every 1 minute (60000 milliseconds)
    const intervalId = setInterval(() => {
      fetchExchangeRates();
    }, 60000);

    // Cleanup interval on component unmount
    return () => clearInterval(intervalId);
  }, []);

  useEffect(() => {
    // Map selectedCoin.name to the corresponding portfolio key
    const coinKeyMap = {
      Bitcoin: 'BTC',
      Ethereum: 'ETH',
      USDC: 'USDC',
    };

    const coinKey = coinKeyMap[selectedCoin.name];

    // Update selectedCoin whenever portfolio changes
    if (portfolio[coinKey]) {
      setSelectedCoin(prevState => ({
        ...prevState,
        price: portfolio[coinKey],
        lastUpdate: new Date().toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit', hour12: true }),
      }));
    }
  }, [portfolio, selectedCoin.name]);

  const handleCoinChange = (coin) => {
    setSelectedCoin({
      name: coin,
      title: coins[coin].title,
      price: portfolio[coin] || 'Loading...',
      lastUpdate: new Date().toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit', hour12: true }),
      icon: coins[coin].icon,
    });
  };

  const coins = {
    Bitcoin: { price: portfolio.BTC || 'Loading...', title: 'Bitcoin', icon: bitcoinLogo },
    Ethereum: { price: portfolio.ETH || 'Loading...', title: 'Ethereum', icon: ethcoinLogo },
    USDC: { price: portfolio.USDC || 'Loading...', title: 'USDC', icon: usdccoinLogo },
  };

  return (
    <div className="app">
      <Ticker />
      <div className="content-box">
        <div className="button-group">
          <button onClick={() => handleCoinChange('Bitcoin')}>Bitcoin</button>
          <button onClick={() => handleCoinChange('Ethereum')}>Ethereum</button>
          <button onClick={() => handleCoinChange('USDC')}>USDC</button>
        </div>
        <BitcoinPrice 
          price={selectedCoin.price} 
          title={selectedCoin.title}
          lastUpdate={selectedCoin.lastUpdate} 
          icon={selectedCoin.icon} 
        />
      </div>
      <footer className="footer">
        <a href="#link1">Link 1</a>
        <a href="#link2">Link 2</a>
        <a href="#link3">Link 3</a>
      </footer>
    </div>
  );
}

export default App;