include .env

run:
	@echo "Running the application..."
	@docker-compose -f docker-compose.yml up --build
	@echo "Application is running."

stop:
	@echo "Stopping the application..."
	@docker-compose -f docker-compose.yml down
	@echo "Application stopped."

clean:
	@echo "Cleaning up..."
	@docker-compose -f docker-compose.yml down --volumes --remove-orphans
	@echo "Cleanup complete."