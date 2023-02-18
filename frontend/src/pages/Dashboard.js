import React from 'react';
import { useSelector } from 'react-redux';

function Dashboard() {
    const auth = useSelector(state => state.auth);

    if(!auth.loggedIn) {
        window.location.replace('/login');
    }

    return (
        <div>
            Dashboard
        </div>
    )
}

export default Dashboard;
