import React, { useState, useEffect } from 'react';

const InvoiceSection = () => {
  const [invoices, setInvoices] = useState([]);
  const [selectedInvoice, setSelectedInvoice] = useState(null);
  const [searchQuery, setSearchQuery] = useState('');
  const [editedProducts, setEditedProducts] = useState([]);

  useEffect(() => {
    fetchInvoices();
  }, []);

  const handleSearch = (query) => {
    fetchInvoices(query);
  };

  const fetchInvoices = (query) => {
    let url = 'http://localhost:8080/invoices';
    if (query) {
      url += `?query=${query}`;
    }
    fetch(url)
      .then(response => response.json())
      .then(data => setInvoices(data))
      .catch(error => console.error('Error fetching invoices:', error));
  };

  const editProduct = (invoice) => {
    setSelectedInvoice(invoice);
    fetchInvoiceProducts(invoice.id);
    // Open modal or navigate to editing page
  };

  const fetchInvoiceProducts = (invoiceId) => {
    fetch(`http://localhost:8080/invoices/${invoiceId}`)
      .then(response => response.json())
      .then(data => setEditedProducts(data.products))
      .catch(error => console.error('Error fetching products:', error));
  };

  const saveEditedProducts = () => {
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
    .catch(error => console.error('Error updating products:', error));
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
      .catch(error => console.error('Error deleting invoice:', error));
    }
  };

  const handleProductChange = (index, newValue) => {
    setEditedProducts(prevProducts => {
      const updatedProducts = [...prevProducts];
      updatedProducts[index] = newValue;
      return updatedProducts;
    });
  };

  return (
    <section id="invoice-section">
      <h2>Pharmacy Invoice</h2>
      <div id="search-bar">
        <input
          type="text"
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          placeholder="Enter search query..."
        />
        <button onClick={() => handleSearch(searchQuery)}>Search</button>
      </div>

      <div>
        <h2>Edit Invoice #{selectedInvoice ? selectedInvoice.id : ''}</h2>
        <ul>
          {editedProducts.map((product, index) => (
            <li key={index}>
              {/* Render input fields for editing each product */}
              <input
                type="text"
                value={product.name}
                onChange={(e) => handleProductChange(index, { ...product, name: e.target.value })}
              />
                <input
                  type="number"
                  value={product.total}
                  onChange={(e) => handleProductChange(index, { ...product, total: parseFloat(e.target.value) })}
                />

              <input
                type="number"
                value={product.quantity}
                onChange={(e) => handleProductChange(index, { ...product, quantity: parseFloat(e.target.value) })}
              />

            <input
              type="number"
              value={product.unit_price}
              onChange={(e) => handleProductChange(index, { ...product, unit_price: parseFloat(e.target.value) })}
            />


              {/* Add other input fields for editing other product properties */}
            </li>
          ))}
        </ul>
        <button onClick={saveEditedProducts}>Save Changes</button>
      </div>

      {invoices.map(invoice => (
        <div key={invoice.id}>
          <div id="customer-info">
            <h3>Customer Information</h3>
            <p><strong>Name:</strong> {invoice.customer_name}</p>
            <p><strong>Email:</strong> {invoice.email}</p>
            <p><strong>Address:</strong> {invoice.address}</p>
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
                      <button onClick={() => editProduct(invoice)}>Edit</button>
                      <button onClick={() => deleteProduct(invoice.id)}>Delete</button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
          <div id="total-amount">
            <p><strong>Total Amount:</strong> {invoice.total_amount}</p>
          </div>
        </div>
      ))}
    </section>
  );
}

export default InvoiceSection;
