import React from 'react';
import { useSelector } from 'react-redux';
import Sidebar from '../../components/Sidebar/Sidebar';
import SearchBar from '../../components/Searchbar/Searchbar';
import './Dashboard.css'

function Dashboard() {
    const auth = useSelector(state => state.auth);

    if(!auth.loggedIn) {
        window.location.replace('/login');
    }

    return (
        <div className='page-container'>
            <Sidebar/>
            <div className='main-content'>
            <SearchBar/>

            </div>
        </div>
    )
}

export default Dashboard;
