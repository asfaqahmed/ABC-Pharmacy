import React, { useState } from 'react';

const Navigation = () => {
  return (
  
  <nav>
        <div className="logo">
            <span><img src='/circle.gif' alt='Abc_logo'/>  </span>
            <span>ABC Pharmacy</span>
        </div> 
        
        <div className='Left'>
        <a href="/">Home</a>
        <a href="/products">Products</a>
        <a href="#">About Us</a>
        <a href="#">Contact Us</a>
        </div>
      </nav>  

  );
}

export default Navigation;
