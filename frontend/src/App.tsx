import React, { useEffect, useState } from 'react';
import logo from './logo.svg';
import './App.css';

type Data = {
  message: string;
};

function App() {
  const [data, setData] = useState<Data | null>(null);

  useEffect(() => {
    fetch('http://localhost:8000/data')
      .then((response) => response.json())
      .then((data) => setData(data))
      .catch((error) => console.error('Error fetching data:', error));
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
        {data && <p>{data.message}</p>}
      </header>
    </div>
  );
}

export default App;