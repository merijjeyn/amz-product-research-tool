import React from 'react';
import './Searchbar.css';

function SearchBar() {
  return (
    <header className="search-bar">
      <form>
        <input type="text" placeholder="" />
        <button type="submit">Search</button>
      </form>
    </header>
  );
}

export default SearchBar;
