import React, { useState, useEffect } from 'react';

function App() {
  const [shops, setShops] = useState([]);

  useEffect(() => {
    fetch('http://localhost:8080/shops') // バックエンドのAPIエンドポイント
      .then(response => response.json())
      .then(data => setShops(data))
      .catch(error => console.error('Error fetching data:', error));
  }, []);

  return (
    <div>
      <h1>Shop Information</h1>
      <ul>
        {shops.map((shop, index) => (
          <li key={index}>
            <strong>Name:</strong> {shop.name},
            <strong>Address:</strong> {shop.address},
            <strong>Tel:</strong> {shop.tel}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;
