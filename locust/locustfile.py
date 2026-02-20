from locust import HttpUser, task, between

class PhonebookUser(HttpUser):
    wait_time = between(1, 5)
    @task(2)
    def list_contacts(self):
        self.client.get("/contacts")
    @task(1)
    def add_contact(self):
        self.client.post("/contacts", json={
            "name": "LoadTest User",
            "phone": "555-0000",
            "email": "test@example.com"
        })
    @task(5)
    def health_check(self):
        self.client.get("/health")