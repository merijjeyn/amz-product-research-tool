import './App.css';

import Homepage from './pages/Homepage';
import Dashboard from './pages/Dashboard';
import Login from './pages/Login/Login';

import Sidebar from './components/Sidebar/Sidebar';
import SearchBar from './components/Searchbar/Searchbar';

import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';


function App() {
  return (
    <Router>
      <div className='App'>
        <Routes>
          <Route path='/' element={<Homepage />}/>
          <Route path='/dashboard' element={<Dashboard />}/>
          <Route path='/login' element={<Login />}/>
        </Routes>
      </div>
    </Router>

    // <div className='page-container'>
    //   <Sidebar/>
    //   <div className='main-content'>
    //     <SearchBar/>

    //   </div>
    // </div>
  );
}

export default App;
