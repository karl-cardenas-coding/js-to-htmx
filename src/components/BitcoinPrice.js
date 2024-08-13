/**
 * Copyright (c) karl-cardenas-coding
 * SPDX-License-Identifier: MIT
 */

import React from 'react';
import './BitcoinPrice.css';

function BitcoinPrice({ price, title, lastUpdate, icon }) {
  return (
    <div className="bitcoin-price-container">
      <div className="heading">
        {title}
      </div>
      <div className="tv-screen">
        <div className="screen-content">
          <img src={icon} alt="Coin Icon" className="coin-icon" />
          <span className="currency">$</span>{price}
        </div>
      </div>
      <div className="last-update">
        Last Update: {lastUpdate}
      </div>
    </div>
  );
}

export default BitcoinPrice;
