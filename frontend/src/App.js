import React, { useState } from "react";
import VoucherForm from "./components/VoucherForm";
import SeatList from "./components/SeatList";
import "./App.css";

function App() {
  const [seats, setSeats] = useState([]);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const handleGenerateVoucher = async (formData) => {
    setLoading(true);
    setError("");  // Reset error message every submit form
    setSeats([]);  // Reset seats every submit form

    try {
      // 1. Check if the voucher is available
      const checkResponse = await fetch("http://localhost:8000/api/check", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          flightNumber: formData.flightNumber,
          date: formData.date,
        }),
      });

      const checkData = await checkResponse.json();

      // 2. If the voucher already exists, display an error message and reset seats.
      if (checkData.data.exists) {
        setError("Voucher sudah pernah dibuat untuk penerbangan ini.");
        setSeats([]);
        setLoading(false);
        return;
      }

      // 3. If the voucher does not exist, proceed to create a new voucher.
      const generateResponse = await fetch("http://localhost:8000/api/generate", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(formData),
      });

      const generateData = await generateResponse.json();

      if (generateData.success) {
        setSeats(generateData.data.seats);
      } else {
        setError("Terjadi kesalahan saat membuat voucher.");
      }
    } catch (err) {
      setError("Terjadi kesalahan dengan server.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="App">
      <h1>Airline Voucher Seat Assignment</h1>
      <VoucherForm onGenerateVoucher={handleGenerateVoucher} />
      {loading && <p>Loading...</p>}
      {error && <p className="error">{error}</p>}
      <SeatList seats={seats} />
    </div>
  );
}

export default App;
