import React, { useState, useEffect } from 'react';
import './App.css';

export default function App() {
  const [res, setRes] = useState('');
  useEffect(() => {
    fetch('http://localhost:3000/hello')
      .then(function (response) {
        response.text()
          .then(resp => {
            setRes(resp)
        })
      });
  }, []);
  return (
    <div className="App">
      <h1>{res}</h1>
    </div>
  );
}