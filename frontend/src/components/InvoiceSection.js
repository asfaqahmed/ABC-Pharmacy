import React, { useState, useEffect } from 'react';

const InvoiceSection = () => {
  const [invoices, setInvoices] = useState([]);

  useEffect(() => {
    fetchInvoices();
  }, []);

  const fetchInvoices = () => {
    fetch('http://localhost:8080/invoices')
      .then(response => response.json())
      .then(data => setInvoices(data))
      .catch(error => console.error('Error fetching invoices:', error));
  };

  const editProduct = () => {
    // editProduct logic here
  };

  const deleteProduct = () => {
    // deleteProduct logic here
  };

  console.log(typeof invoices);
  console.log(Array.isArray(invoices));

  return (

    <section id="invoice-section">
      <h2>Pharmacy Invoice</h2>
      {/* {invoices.map(invoice => (
        <div key={invoices.id}>
          <div id="customer-info">
            <h3>Customer Information</h3>
            <p><strong>Name:</strong> {invoices.name}</p>
            <p><strong>Email:</strong> {invoices.email}</p>
            <p><strong>Address:</strong> {invoices.address}</p>
          </div> */}
      


          {/* <div id="product-details">
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
                      <button onClick={editProduct}>Edit</button>
                      <button onClick={deleteProduct}>Delete</button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>

          <div id="total-amount">
            <p><strong>Total Amount:</strong> {invoices.total_amount}</p>
          </div>
        </div> */}
      {/* ))} */}
      

    </section>
  );
}

export default InvoiceSection;
