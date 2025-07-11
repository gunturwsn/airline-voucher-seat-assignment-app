import axios from "axios";

const API_BASE_URL = "http://localhost:8000/api";

export const checkVoucher = async (flightNumber, date) => {
  const response = await axios.post(`${API_BASE_URL}/check`, {
    flightNumber,
    date,
  });
  return response.data;
};

export const generateVoucher = async (formData) => {
  const response = await axios.post(`${API_BASE_URL}/generate`, formData);
  return response.data;
};
