Test Bidding Service
	Scenario - Bidding Sever receives HTTP request
		Given request handler 
		Then service receives request 
			When request contains AdPlacmentID 
				Then service should respond with AdOpbject 
					Then AdObject should contain AdID
					And  AdObject should contain BidPrice ( random for now )
						Then BidPrice should be in cents ( avoid floats for currency !!! ) 
					And StatusOK (200)
				When service not interested in the spot
					Then return StatusNoContent (204)
			But the request doe's not contains AdPlacementID
				Then Status Forbidden 
			
			
	


