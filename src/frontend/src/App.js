import React from 'react';
import "./App.css"
import { BrowserRouter as Router, Routes, Route, Switch} from 'react-router-dom';
import SideBar from './components/SideBar/SideBar';
import MainGPT from './pages/MainGPT/MainGPT';
import FirstPage from './pages/FirstPage/FirstPage'
import About from './pages/About/About'




  
function App() {
  return (

      
    <Router>
      <Routes>
          <Route path='/' active element={<FirstPage/>} />
          <Route path='/MainGPT' element={<MainGPT/>} />
          <Route path='/About' element={<About/>} />          
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
