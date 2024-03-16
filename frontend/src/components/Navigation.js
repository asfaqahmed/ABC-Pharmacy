import React, { useState } from 'react';

const Navigation = () => {
  const [loading, setLoading] = useState(true);
  return (
    <nav>
      
    <div className="logo">
       <span> <img src='/circle.gif' alt='Chemistry'/> </span> <span>ABC Pharmacy</span>
    </div> 
    
    <div className='Left'>
    <a href="/">Home</a>
    <a href="#">Products</a>
    <a href="#">About Us</a>
    <a href="#">Contact Us</a>
    </div>
  </nav>

  );
}

export default Navigation;
