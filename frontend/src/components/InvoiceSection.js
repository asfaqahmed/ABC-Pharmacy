import React, { useState, useEffect } from 'react';

const InvoiceSection = () => {
  const [invoices, setInvoices] = useState([]);
  const [selectedInvoice, setSelectedInvoice] = useState(null);
  const [searchQuery, setSearchQuery] = useState('');
  const [editedProducts, setEditedProducts] = useState([]);
  const [isEditing, setIsEditing] = useState(false);
  const [editedProduct, setEditedProduct] = useState({});
  const [saveMessage, setSaveMessage] = useState('');
  const [errorMessage, setErrorMessage] = useState('');

  useEffect(() => {
    fetchInvoices();
  }, []);

  const handleSearch = () => {
    if (!searchQuery || !invoices) {
      return invoices;
    } else {
      return invoices.filter(invoice =>
        invoice.customer_name.toLowerCase().includes(searchQuery.toLowerCase())
      );
    }
  };

  const fetchInvoices = (query) => {
    let url = 'http://localhost:8080/invoices';
    if (query) {
      url += `?query=${query}`;
    }
    fetch(url)
      .then(response => response.json())
      .then(data => {
        setInvoices(data);
        console.log('Fetched invoices:', data); // Log fetched invoices
      })
      .catch(error => {
        console.error('Error fetching invoices:', error);
        setErrorMessage('Error fetching invoices. Please try again later.'); // Set error message
      });
  };

  const editProduct = (invoice) => {
    setSelectedInvoice(invoice);
    fetchInvoiceProducts(invoice.id);
  };

  const fetchInvoiceProducts = (invoiceId) => {
    fetch(`http://localhost:8080/invoices/${invoiceId}`)
      .then(response => response.json())
      .then(data => {
        setEditedProducts(data.products);
        console.log('Fetched invoice products:', data.products); // Log fetched invoice products
      })
      .catch(error => {
        console.error('Error fetching products:', error);
        setErrorMessage('Error fetching products. Please try again later.'); // Set error message
      });
  };

  const saveEditedProducts = () => {
    if (editedProducts.length === 0) {
      console.log('No changes to save');
      return;
    }
    
    fetch(`http://localhost:8080/invoices/${selectedInvoice.id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ products: editedProducts })
      
    })
    .then(response => {
      if (response.ok) {
        console.log('Products updated successfully');
        // Redirect to the invoice details page or update the state accordingly
      } else {
        throw new Error('Failed to update products');
      }
    })
    .then(() => {
      setSaveMessage('Products updated successfully');
    })
    .catch(error => {
      console.error('Error updating products:', error);
      setErrorMessage('Error updating products. Please try again later.'); // Set error message
    });
  };

  const handleProductChange = (index, newValue) => {
    setEditedProducts(prevProducts => {
      const updatedProducts = [...prevProducts];
      updatedProducts[index] = newValue;
      return updatedProducts;
    });
  };

  const handleEditChange = (propertyName, value) => {
    setEditedProduct({ ...editedProduct, [propertyName]: value });
  };

  const deleteProduct = (invoiceId) => {
    if (window.confirm('Are you sure you want to delete this invoice?')) {
      fetch(`http://localhost:8080/invoices/${invoiceId}`, {
        method: 'DELETE'
      })
      .then(response => {
        if (response.ok) {
          setInvoices(prevInvoices => prevInvoices.filter(invoice => invoice.id !== invoiceId));
          console.log('Invoice deleted successfully');
        } else {
          throw new Error('Failed to delete invoice');
        }
      })
      .catch(error => {
        console.error('Error deleting invoice:', error);
        setErrorMessage('Error deleting invoice. Please try again later.'); // Set error message
      });
    }
  };

  return (
    <section>
      <div className="search-bar" id="search-bar">
        <input
          type="text"
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          placeholder="Enter search query..."
        />
        <button className="bl-button" onClick={handleSearch}>Search</button>
      </div>
      <div>
        {errorMessage && <p>Error: {errorMessage}</p>} {/* Display error message */}
        {saveMessage && <p>{saveMessage}</p>}
        <h2>Edit Invoice #{selectedInvoice ? selectedInvoice.id : ''}</h2>
        <ul>
          {editedProducts.map((product, index) => (
            <li key={index}>
              <label>Name :</label>
              <input
                type="text"
                value={product.name}
                onChange={(e) => handleProductChange(index, { ...product, name: e.target.value })}
              />

              <label>Quantity :</label>
              <input
                type="number"
                value={product.quantity}
                onChange={(e) => handleProductChange(index, { ...product, quantity: parseFloat(e.target.value) })}
              />

              <label>Unit Price :</label>
              <input
                type="number"
                value={product.unit_price}
                onChange={(e) => handleProductChange(index, { ...product, unit_price: parseFloat(e.target.value) })}
              />

              <label>Total :</label>
              <input
                type="number"
                value={product.total}
                onChange={(e) => handleProductChange(index, { ...product, total: parseFloat(e.target.value) })}
              />
            </li>
          ))}
        </ul>
        <button onClick={saveEditedProducts}>Save Changes</button>
      </div>

      
      {invoices.map(invoice => (
  <div key={invoice.id}>
    <div id="customer-info">
      <h3>Customer Information</h3>
      <p><strong>Name:</strong> {invoice.customer_name || "John Doe"}</p>
      <p><strong>Email:</strong> {invoice.email || "John@example.com"}</p>
      <p><strong>Address:</strong> {invoice.address || "123 Main St, City"}</p>
    </div>
    
    <div id="product-details">
      <h3>Product Details</h3>
      <table>
        <thead>
          <tr>
            <th>Product</th>
            <th>Quantity</th>
            <th>Price per Unit</th>
            <th>Total</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          {invoice.products.map(product => (
            <tr key={product.id}>
              <td>{product.name}</td>
              <td>{product.quantity}</td>
              <td>{product.unit_price}</td>
              <td>{product.total}</td>
              <td>
                <button className="bl-button" onClick={() => editProduct(invoice)}>Edit</button>
                <button className="bl-button" onClick={() => deleteProduct(invoice.id)}>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
    
    <div id="total-amount">
      <p><strong>Total Amount:</strong> {invoice.products.reduce((total, product) => total + product.total, 0)}</p>
    </div>
  </div>
))}

    </section>
  );
}

export default InvoiceSection;
