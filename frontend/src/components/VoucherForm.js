import React, { useState } from "react";

function VoucherForm({ onGenerateVoucher }) {
  const [formData, setFormData] = useState({
    name: "",
    id: "",
    flightNumber: "",
    date: "",
    aircraft: "ATR",
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onGenerateVoucher(formData);
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label>Name:</label>
        <input
          type="text"
          name="name"
          value={formData.name}
          onChange={handleChange}
          required
        />
      </div>
      <div>
        <label>ID:</label>
        <input
          type="text"
          name="id"
          value={formData.id}
          onChange={handleChange}
          required
        />
      </div>
      <div>
        <label>Flight Number:</label>
        <input
          type="text"
          name="flightNumber"
          value={formData.flightNumber}
          onChange={handleChange}
          required
        />
      </div>
      <div>
        <label>Flight Date:</label>
        <input
          type="date"
          name="date"
          value={formData.date}
          onChange={handleChange}
          required
        />
      </div>
      <div>
        <label>Aircraft Type:</label>
        <select
          name="aircraft"
          value={formData.aircraft}
          onChange={handleChange}
        >
          <option value="ATR">ATR</option>
          <option value="Airbus 320">Airbus 320</option>
          <option value="Boeing 737 Max">Boeing 737 Max</option>
        </select>
      </div>
      <button type="submit">Generate Vouchers</button>
    </form>
  );
}

export default VoucherForm;
