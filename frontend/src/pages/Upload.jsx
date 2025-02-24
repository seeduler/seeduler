import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import axios from '../config/axios'
import '../styles/Upload.css'

function Upload() {
  const [jsonData, setJsonData] = useState('')
  const [error, setError] = useState('')
  const [credentials, setCredentials] = useState(null)
  const navigate = useNavigate()

  const handleSubmit = async (e) => {
    e.preventDefault()
    try {
      const data = JSON.parse(jsonData)
      const response = await axios.post('/halls/upload-data', data)
      setCredentials(response.data)
    } catch (err) {
      setError(err instanceof SyntaxError ? 'Invalid JSON format' : err.response?.data || 'Upload failed')
    }
  }

  return (
    <div className="upload-container">
      <div className="upload-card">
        {!credentials ? (
          <>
            <h1>Initialize System</h1>
            <p className="instruction">Paste your JSON data below:</p>
            <form onSubmit={handleSubmit}>
              <textarea
                value={jsonData}
                onChange={(e) => setJsonData(e.target.value)}
                placeholder="Paste your JSON here..."
                rows={10}
              />
              {error && <div className="error-message">{error}</div>}
              <button type="submit" className="submit-btn">
                Initialize
              </button>
            </form>
          </>
        ) : (
          <div className="credentials-container">
            <h2>Generated Credentials</h2>
            <div className="warning-banner">
              ⚠️ Save these credentials now. You won't see them again!
            </div>
            <div className="credentials-list">
              {credentials.map(user => (
                <div key={user.hall_id} className="credential-item">
                  <h3>{user.username}</h3>
                  <p>Password: {user.password}</p>
                </div>
              ))}
            </div>
            <button 
              onClick={() => navigate('/login')} 
              className="continue-btn"
            >
              Continue to Login
            </button>
          </div>
        )}
      </div>
    </div>
  )
}

export default Upload 