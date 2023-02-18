import React from 'react';
import './SidebarEntry.css'

function SidebarEntry(props) {
    return(
        <div className='sidebar-entry-container'>
            {props.children}
        </div>  
    );
}

export default SidebarEntry;