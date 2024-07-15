import React, { useState, useEffect } from 'react';
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';

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
      <MapContainer center={[shops[0]?.lat || 0, shops[0]?.lng || 0]} zoom={10} style={{ height: '400px', width: '800px' }}>
        <TileLayer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          attribution="&copy; <a href='http://osm.org/copyright'>OpenStreetMap</a> contributors"
        />
        {shops.map((shop, index) => (
          <Marker key={index} position={[shop.lat, shop.lng]}>
            <Popup>
              <strong>Name:</strong> {shop.name}<br />
              <strong>Address:</strong> {shop.address}<br />
              <img src={shop.photo_url} alt={`${shop.name}`} style={{ width: '100%', height: 'auto' }} />
            </Popup>
          </Marker>
        ))}
      </MapContainer>
      <ul>
        {shops.map((shop, index) => (
          <li key={index}>
            <strong>店舗名:</strong> {shop.name}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;
