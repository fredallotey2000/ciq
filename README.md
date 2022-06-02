# CIQs log query solution

# OPENING THE SOLUTION 
1. unzip the project or clone a copy from https://github.com/fredallotey2000/ciq.git
2. open VScode
3. go to --> file --> open folder --> browse to the location of the unzipped folder
4. open a new terminal 


# TESTING THE SOLUTION
1. run go test -v ./...

# BENCHMARD THE SOLUTION
1. run go test ./...  -run=Benchmark -bench=.

# TO RUN THE SOLUTION 
1. Start a new terminal 
2. At the termnial --> go run main.go <OPTIONS>
OR using the build ./main <OPTIONS> 
3. Where <oPTIONS> are as follows

<oPTIONS>
1. -d Operation date of the log.ie. the the log was registered
2. -o Operation name. ie. either upload or download
4. -u Username of user who performed the operation
5. -s Download or Upload size registered by the log
5. -u all this flag specifies all users.eg. usage ./main -u all (outcome will be number of users)

# Examples
1. go run main.go -u rosannaM -d Apr-19-2020 -o download
2. go run main.go -u rosannaM -d apr-19-2020 -o upload
4. go run main.go -u rosannaM -d apr-19-2020
5. go run main.go -s 50
6. go run main.go -u all (for all users)



