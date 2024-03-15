import React, { Fragment } from 'react';
import { useState } from 'react';

const CrudSection = () => {
  const [productName, setProductName] = useState('');
  const [price, setPrice] = useState('');
  const [itemCategory, setItemCategory] = useState('');

  const [customerName, setCustomerName] = useState('');
  const [mobileNo, setMobileNo] = useState('');
  const [email, setEmail] = useState('');
  const [address, setAddress] = useState('');
  const [billingType, setBillingType] = useState('');
  

  const addProduct = () => {
    // Your addProduct logic here
    const apiUrl = 'http://localhost:8080/items';

    const priceFloat = parseFloat(price);

    const payload = {
      name: productName,
      unit_price: priceFloat,
      item_category : itemCategory
    };

    fetch(apiUrl, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(payload),
  })
    .then(response => response.json())
    .then(data => {
      console.log('Product added successfully:', data);
      // Add any additional logic or state updates as needed
    })
    .catch(error => {
      console.error('Error adding product:', error);
    });
};
  const addCustomer = () => {
    const apiUrl = 'http://localhost:8080/customers';

    const payload = {
      name: customerName,
      mobile_no: mobileNo,
      email : email,
      address : address,
      billingType : billingType

    };

    fetch(apiUrl, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(payload),
  })
    .then(response => response.json())
    .then(data => {
      console.log('Product added successfully:', data);
      // Add any additional logic or state updates as needed
    })
    .catch(error => {
      console.error('Error adding product:', error);
    });
  };

  return (

<section id="crud-section">
      <div className="crud-container">
        <div className="crud-section">
          <h2>Manage Products</h2>
          <form id="product-form">
            {/* Product form fields */}
            <label htmlFor="productName">Product Name:</label>
        <input
          type="text"
          id="productName"
          name="productName"
          value={productName}
          onChange={(e) => setProductName(e.target.value)}
          required
        />

      <label htmlFor="price">Price per Unit:</label>
        <input
          type="number"
          id="price"
          name="price"
          value={price}
          onChange={(e) => setPrice(e.target.value)}
          required
        />

        <label htmlFor="itemCategory">itemCategory:</label>
        <input
          type="text"
          id="itemCategory"
          name="itemCategory"
          value={itemCategory}
          onChange={(e) => setItemCategory(e.target.value)}
          required
        />
            <button type="button" onClick={addProduct}>
              Add Product
            </button>
          </form>
        </div>

    {/* Customer  fields */}
        <div className="crud-section">
          <h2>Manage Customers</h2>
          <form id="customer-form">
      <label htmlFor="CustomerForm">Customer Name:</label>
        <input
          type="text"
          id="customerName"
          name="customerName"
          value={customerName}
          onChange={(e) => setCustomerName(e.target.value)}
          required
        />

        <label htmlFor="Mobile">Mobile NO:</label>
        <input
          type="number"
          id="mobileNo"
          name="mobileNo"
          value={mobileNo}
          onChange={(e) => setMobileNo(e.target.value)}
          required
        />

        <label htmlFor="Email">Email</label>
        <input
          type="Email"
          id="email"
          name="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />

      <label htmlFor="Address">Address </label>
        <input
          type="textBox"
          id="address"
          name="address"
          value={address}
          onChange={(e) => setAddress(e.target.value)}
          required
        />

      <label htmlFor="billingType">BillingType</label>
        <input
          type="Choice"
          id="billingType"
          name="billingType"
          value={billingType}
          onChange={(e) => setBillingType(e.target.value)}
          required
        />

            <button type="button" onClick={addCustomer}>
              Add Customer
            </button>
          </form>
          <section></section>
        </div>
      </div>
    </section>

  );
}

export default CrudSection;
