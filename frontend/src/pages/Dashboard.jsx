import { useState, useEffect } from 'react'
import axios from 'axios'
import { useAuth } from '../contexts/AuthContext'

function Dashboard() {
  const [events, setEvents] = useState([])
  const { token } = useAuth()

  useEffect(() => {
    const fetchEvents = async () => {
      const response = await axios.get('/halls/with-events', {
        headers: { Authorization: `Bearer ${token}` }
      })
      setEvents(response.data)
    }
    fetchEvents()
  }, [token])

  const handleMarkComplete = async (eventId) => {
    await axios.post('/events/mark-completed', 
      { event_id: eventId },
      { headers: { Authorization: `Bearer ${token}` }}
    )
    // Refresh events
  }

  const handleAddDelay = async (eventId, delay) => {
    await axios.post('/events/add-delay',
      { event_id: eventId, delay },
      { headers: { Authorization: `Bearer ${token}` }}
    )
    // Refresh events
  }

  return (
    <div className="dashboard">
      {events.map(hall => (
        <div key={hall.hall.id} className="hall-section">
          <h2>{hall.hall.name}</h2>
          <p>Current Delay: {hall.hall.delayed_time} minutes</p>
          
          <div className="events-list">
            {hall.events.map(event => (
              <div key={event.id} className="event-card">
                <h3>{event.title}</h3>
                <p>Start: {new Date(event.start_time).toLocaleString()}</p>
                <p>End: {new Date(event.end_time).toLocaleString()}</p>
                <button
                  onClick={() => handleMarkComplete(event.id)}
                  disabled={event.is_completed}
                >
                  {event.is_completed ? 'Completed' : 'Mark Complete'}
                </button>
                <button onClick={() => handleAddDelay(event.id, 5)}>
                  Add 5min Delay
                </button>
              </div>
            ))}
          </div>
        </div>
      ))}
    </div>
  )
}

export default Dashboard 