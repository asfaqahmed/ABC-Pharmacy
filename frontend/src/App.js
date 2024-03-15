import React from 'react';
import Navigation from './components/Navigation';
import InvoiceSection from './components/InvoiceSection';
import CrudSection from './components/CrudSection';
import Footer from './components/Footer';
import './App.css';

function App() {
  return (
    <div className="app">
      <Navigation />
      <InvoiceSection />
      <CrudSection />
      <Footer />
    </div>
  );
}

export default App;
