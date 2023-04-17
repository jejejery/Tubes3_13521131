import React from 'react';
import "./App.css"
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import SideBar from './components/SideBar/SideBar';
import MainGPT from './pages/MainGPT/MainGPT';



  
function App() {
  return (

      
    <Router>
      <Routes>
        {/* <Route path='/' active element={<Home/>} /> */}
        <Route path='/MainGPT' element={<MainGPT/>} />
      </Routes>
    </Router>
  
  );
}
  
export default App;

// import logo from './logo.svg';
// import './App.css';

// function App() {
//   return (
//     <div className="App">
//       <header className="App-header">
//         <img src={logo} className="App-logo" alt="logo" />
//         <p>
//           Edit <code>src/App.js</code> and save to reload.
//         </p>
//         <a
//           className="App-link"
//           href="https://reactjs.org"
//           target="_blank"
//           rel="noopener noreferrer"
//         >
//           Learn React
//         </a>
//       </header>
//     </div>
//   );
// }

// export default App;
