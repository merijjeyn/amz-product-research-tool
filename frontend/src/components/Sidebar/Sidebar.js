import React from 'react';
import './Sidebar.css';
import SidebarEntryContent from './SidebarEntry/SidebarEntryContent';

function Sidebar() {
  return (
    <aside className="sidebar">
        <SidebarEntryContent />
    </aside>
  );
}

export default Sidebar;
