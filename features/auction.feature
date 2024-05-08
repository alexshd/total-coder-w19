Test Auction Service
	Given multiple bidding services
		The Auction service sends concurrent requests 
			When the response StatusCodeOK (200) 
				Then accumulate valid bids
			When response StatusCode != 200
				Or Bidding Service response time over 200ms
					Then Auction should  not accept the bid from that Bidding Service 
			When FanOut finished
				Then return the Max BidPrice  		
				
			
Test Auction Service Client Exposed API
	Given publicly exposed API 
		On external client request
			When contains AdPlacementID
				Then "FanOut" ( optimize ) client request
					When no eligible bids are received 
						Then return StatusCodeNoContent ( 204 ) to the client
					When On Success 
						Then return Max BidPrice 
