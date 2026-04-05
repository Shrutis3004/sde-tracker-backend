package database

func cnSheet() sheetData {
	return sheetData{
		Name:        "CN Sheet",
		Slug:        "cn-sheet",
		Description: "Most Asked Computer Networks Interview Questions",
		Topics: []topicData{
			{Name: "Introduction to CN", Problems: []problemData{
				{"What is Network", "https://www.geeksforgeeks.org/basics-computer-networking/", "Easy"},
				{"Types of Network (LAN, WAN, MAN)", "https://www.geeksforgeeks.org/types-of-area-networks-lan-man-and-wan/", "Easy"},
				{"Hub, Switch, and Router", "https://www.geeksforgeeks.org/difference-between-hub-switch-and-router/", "Easy"},
				{"Network Topology and its Types", "https://www.geeksforgeeks.org/types-of-network-topology/", "Easy"},
				{"Nodes and Links", "https://www.geeksforgeeks.org/network-devices-hub-repeater-bridge-switch-router-gateways/", "Easy"},
			}},
			{Name: "OSI Model", Problems: []problemData{
				{"OSI Model Overview", "https://www.geeksforgeeks.org/open-systems-interconnection-model-osi/", "Easy"},
				{"Layers of OSI Model", "https://www.geeksforgeeks.org/layers-of-osi-model/", "Medium"},
			}},
			{Name: "TCP/IP Model", Problems: []problemData{
				{"TCP/IP Protocol Suite", "https://www.geeksforgeeks.org/tcp-ip-model/", "Medium"},
				{"Different Layers of TCP/IP Model", "https://www.geeksforgeeks.org/tcp-ip-model/", "Medium"},
				{"Difference between OSI and TCP/IP Model", "https://www.geeksforgeeks.org/difference-between-osi-and-tcp-ip-reference-model/", "Easy"},
			}},
			{Name: "Data Link Layer", Problems: []problemData{
				{"Data Link Layer and its Functions", "https://www.geeksforgeeks.org/data-link-layer/", "Medium"},
				{"Error Detection and Correction", "https://www.geeksforgeeks.org/error-detection-in-computer-networks/", "Medium"},
				{"Data Link Protocols", "https://www.geeksforgeeks.org/data-link-layer-protocols/", "Medium"},
				{"Multiple Access Protocols", "https://www.geeksforgeeks.org/multiple-access-protocols-in-computer-network/", "Medium"},
				{"MAC Address", "https://www.geeksforgeeks.org/mac-address-in-computer-network/", "Easy"},
				{"Channel Allocation Problem", "https://www.geeksforgeeks.org/channel-allocation-problem-in-computer-network/", "Medium"},
				{"Data Link Layer Switching", "https://www.geeksforgeeks.org/data-link-layer-switching/", "Medium"},
				{"Ethernet LANs", "https://www.geeksforgeeks.org/local-area-network-lan-technologies/", "Medium"},
			}},
			{Name: "Network Layer", Problems: []problemData{
				{"Congestion Control Algorithms", "https://www.geeksforgeeks.org/congestion-control-in-computer-networks/", "Hard"},
				{"IP Addressing and Subnetting", "https://www.geeksforgeeks.org/introduction-of-classful-ip-addressing/", "Medium"},
				{"IPv4 and IPv6 Protocols", "https://www.geeksforgeeks.org/differences-between-ipv4-and-ipv6/", "Medium"},
				{"ICMP", "https://www.geeksforgeeks.org/internet-control-message-protocol-icmp/", "Medium"},
				{"Fragmentation", "https://www.geeksforgeeks.org/fragmentation-at-network-layer/", "Medium"},
				{"Tunneling", "https://www.geeksforgeeks.org/computer-network-tunneling/", "Medium"},
				{"Subnet", "https://www.geeksforgeeks.org/introduction-of-subnetting/", "Medium"},
				{"Gateway and Router", "https://www.geeksforgeeks.org/difference-between-router-and-gateway/", "Easy"},
				{"Count to Infinity Problem", "https://www.geeksforgeeks.org/count-to-infinity-problem-in-distance-vector-routing/", "Hard"},
			}},
			{Name: "Transport Layer", Problems: []problemData{
				{"Elements of Transport Protocol", "https://www.geeksforgeeks.org/transport-layer-responsibilities/", "Medium"},
				{"TCP Protocol", "https://www.geeksforgeeks.org/what-is-transmission-control-protocol-tcp/", "Medium"},
				{"UDP Protocol", "https://www.geeksforgeeks.org/user-datagram-protocol-udp/", "Medium"},
				{"TCP vs UDP", "https://www.geeksforgeeks.org/differences-between-tcp-and-udp/", "Easy"},
				{"3-Way Handshaking", "https://www.geeksforgeeks.org/tcp-3-way-handshake-process/", "Medium"},
				{"Congestion Control Mechanisms", "https://www.geeksforgeeks.org/tcp-congestion-control/", "Hard"},
			}},
			{Name: "Application Layer", Problems: []problemData{
				{"Domain Name System (DNS)", "https://www.geeksforgeeks.org/domain-name-system-dns-in-application-layer/", "Medium"},
				{"World Wide Web (WWW)", "https://www.geeksforgeeks.org/world-wide-web-www/", "Easy"},
				{"HTTP Protocol", "https://www.geeksforgeeks.org/http-full-form/", "Medium"},
				{"FTP Protocol", "https://www.geeksforgeeks.org/file-transfer-protocol-ftp-in-application-layer/", "Medium"},
				{"SMTP Protocol", "https://www.geeksforgeeks.org/simple-mail-transfer-protocol-smtp/", "Medium"},
				{"DHCP Protocol", "https://www.geeksforgeeks.org/dynamic-host-configuration-protocol-dhcp/", "Medium"},
			}},
			{Name: "Network Security", Problems: []problemData{
				{"What is Network Security", "https://www.geeksforgeeks.org/network-security/", "Easy"},
				{"Cryptography", "https://www.geeksforgeeks.org/cryptography-and-its-types/", "Medium"},
				{"Types of Encryption (RSA, DES, AES)", "https://www.geeksforgeeks.org/types-of-encryption-algorithms/", "Medium"},
				{"Hash Functions", "https://www.geeksforgeeks.org/hash-functions-and-list-of-all-hash-functions/", "Medium"},
				{"Digital Signatures", "https://www.geeksforgeeks.org/digital-signatures-certificates/", "Medium"},
				{"Network Security Protocols", "https://www.geeksforgeeks.org/network-security-protocols/", "Medium"},
				{"Virtual Private Networks", "https://www.geeksforgeeks.org/virtual-private-network-vpn-introduction/", "Medium"},
				{"Network Attacks", "https://www.geeksforgeeks.org/types-of-network-attacks/", "Medium"},
			}},
		},
	}
}

func dbmsSheet() sheetData {
	return sheetData{
		Name:        "DBMS Sheet",
		Slug:        "dbms-sheet",
		Description: "Most Asked DBMS Interview Questions",
		Topics: []topicData{
			{Name: "Introduction to DBMS", Problems: []problemData{
				{"Data, Database and File System", "https://www.geeksforgeeks.org/difference-between-file-system-and-dbms/", "Easy"},
				{"Types of Database", "https://www.geeksforgeeks.org/types-of-databases/", "Easy"},
				{"What is DBMS and its Applications", "https://www.geeksforgeeks.org/dbms-introduction-database-management-system-set-1/", "Easy"},
				{"Data Abstraction", "https://www.geeksforgeeks.org/data-abstraction-and-data-independence/", "Easy"},
				{"3-Tier Architecture", "https://www.geeksforgeeks.org/introduction-of-3-tier-architecture-in-dbms-set-2/", "Medium"},
			}},
			{Name: "Data Models", Problems: []problemData{
				{"Data Models and Types", "https://www.geeksforgeeks.org/data-models-in-dbms/", "Easy"},
				{"ER Model and Components", "https://www.geeksforgeeks.org/introduction-of-er-model/", "Medium"},
				{"Entity-Relationship Diagrams", "https://www.geeksforgeeks.org/er-model-enhanced-er-model/", "Medium"},
				{"Relational Model", "https://www.geeksforgeeks.org/relational-model-in-dbms/", "Medium"},
			}},
			{Name: "RDBMS", Problems: []problemData{
				{"What is RDBMS", "https://www.geeksforgeeks.org/rdbms-full-form/", "Easy"},
				{"Essential Components of Table", "https://www.geeksforgeeks.org/database-management-system-relational-algebra/", "Easy"},
				{"Intension and Extension in Database", "https://www.geeksforgeeks.org/database-schemas/", "Medium"},
				{"Keys in DBMS", "https://www.geeksforgeeks.org/types-of-keys-in-relational-model-candidate-super-primary-alternate-and-foreign/", "Medium"},
				{"Normalization and Types (1NF, 2NF, 3NF)", "https://www.geeksforgeeks.org/normal-forms-in-dbms/", "Medium"},
				{"Denormalization", "https://www.geeksforgeeks.org/denormalization-in-databases/", "Medium"},
			}},
			{Name: "SQL", Problems: []problemData{
				{"SQL Commands (INSERT, UPDATE, DELETE, SELECT)", "https://www.geeksforgeeks.org/sql-ddl-dql-dml-dcl-tcl-commands/", "Medium"},
				{"Operators and Aggregation in SQL", "https://www.geeksforgeeks.org/sql-operators/", "Medium"},
				{"SQL Sub Queries", "https://www.geeksforgeeks.org/sql-subquery/", "Medium"},
			}},
			{Name: "Transactions & Concurrency", Problems: []problemData{
				{"Schedule (Serial, Parallel)", "https://www.geeksforgeeks.org/transaction-schedules-in-dbms/", "Medium"},
				{"Isolation Levels and Types", "https://www.geeksforgeeks.org/transaction-isolation-levels-dbms/", "Medium"},
				{"Serializability and Concurrency Control", "https://www.geeksforgeeks.org/concurrency-control-in-dbms/", "Hard"},
				{"Locking Protocols", "https://www.geeksforgeeks.org/lock-based-concurrency-control-protocol-in-dbms/", "Hard"},
			}},
			{Name: "Query Optimization", Problems: []problemData{
				{"Techniques for Optimizing SQL Queries", "https://www.geeksforgeeks.org/sql-query-optimization/", "Medium"},
				{"Indexing and its Types", "https://www.geeksforgeeks.org/indexing-in-databases-set-1/", "Medium"},
				{"B and B+ Trees", "https://www.geeksforgeeks.org/introduction-of-b-tree-2/", "Hard"},
			}},
			{Name: "Scalability & Performance", Problems: []problemData{
				{"Vertical and Horizontal Scaling", "https://www.geeksforgeeks.org/horizontal-and-vertical-scaling-in-databases/", "Medium"},
				{"Sharding", "https://www.geeksforgeeks.org/database-sharding-a-system-design-concept/", "Medium"},
			}},
			{Name: "Data Integrity & Security", Problems: []problemData{
				{"RBAC (Role-Based Access Control)", "https://www.geeksforgeeks.org/role-based-access-control/", "Medium"},
				{"Encryption", "https://www.geeksforgeeks.org/types-of-encryption-algorithms/", "Medium"},
			}},
		},
	}
}

func osSheet() sheetData {
	return sheetData{
		Name:        "OS Sheet",
		Slug:        "os-sheet",
		Description: "Most Asked Operating System Interview Questions",
		Topics: []topicData{
			{Name: "Introduction to OS", Problems: []problemData{
				{"OS and Main Functions", "https://www.geeksforgeeks.org/introduction-of-operating-system-set-1/", "Easy"},
				{"Types of OS", "https://www.geeksforgeeks.org/types-of-operating-systems/", "Easy"},
				{"Batch OS", "https://www.geeksforgeeks.org/batch-processing-operating-system/", "Easy"},
				{"Multiprogramming and Multitasking OS", "https://www.geeksforgeeks.org/difference-between-multiprogramming-and-multitasking/", "Easy"},
				{"Kernel and User Modes", "https://www.geeksforgeeks.org/difference-between-user-mode-and-kernel-mode/", "Medium"},
				{"Process and its States", "https://www.geeksforgeeks.org/states-of-a-process-in-operating-systems/", "Medium"},
				{"System Calls in OS", "https://www.geeksforgeeks.org/introduction-of-system-call/", "Medium"},
			}},
			{Name: "Process Management", Problems: []problemData{
				{"Starvation and Aging", "https://www.geeksforgeeks.org/starvation-and-aging-in-operating-systems/", "Medium"},
				{"Process Synchronization", "https://www.geeksforgeeks.org/introduction-of-process-synchronization/", "Medium"},
				{"Semaphores and Types", "https://www.geeksforgeeks.org/semaphores-in-process-synchronization/", "Medium"},
				{"Deadlock and Prevention Methods", "https://www.geeksforgeeks.org/introduction-of-deadlock-in-operating-system/", "Hard"},
				{"Banker's Algorithm", "https://www.geeksforgeeks.org/bankers-algorithm-in-operating-system-2/", "Hard"},
			}},
			{Name: "Memory Management", Problems: []problemData{
				{"Memory Management Techniques", "https://www.geeksforgeeks.org/memory-management-in-operating-system/", "Medium"},
				{"Virtual Memory", "https://www.geeksforgeeks.org/virtual-memory-in-operating-system/", "Medium"},
				{"Dynamic Binding", "https://www.geeksforgeeks.org/static-and-dynamic-binding/", "Medium"},
				{"Page Fault", "https://www.geeksforgeeks.org/page-fault-handling-in-operating-system/", "Medium"},
				{"Belady's Anomaly", "https://www.geeksforgeeks.org/beladys-anomaly-in-page-replacement-algorithms/", "Medium"},
				{"Thrashing", "https://www.geeksforgeeks.org/techniques-to-handle-thrashing/", "Medium"},
				{"Cache", "https://www.geeksforgeeks.org/cache-memory-in-computer-organization/", "Medium"},
				{"Direct and Associative Mapping", "https://www.geeksforgeeks.org/cache-mapping/", "Hard"},
			}},
			{Name: "File Systems", Problems: []problemData{
				{"File System and Components", "https://www.geeksforgeeks.org/file-systems-in-operating-system/", "Medium"},
				{"Types of File System", "https://www.geeksforgeeks.org/different-types-of-file-systems/", "Medium"},
				{"File Allocation and Deallocation", "https://www.geeksforgeeks.org/file-allocation-methods/", "Medium"},
				{"Fragmentation", "https://www.geeksforgeeks.org/difference-between-internal-and-external-fragmentation/", "Medium"},
			}},
			{Name: "I/O Systems", Problems: []problemData{
				{"Blocking vs Non-Blocking", "https://www.geeksforgeeks.org/difference-between-blocking-and-non-blocking/", "Easy"},
				{"Synchronous vs Asynchronous", "https://www.geeksforgeeks.org/synchronous-and-asynchronous-in-operating-system/", "Easy"},
				{"Spooling", "https://www.geeksforgeeks.org/what-exactly-spooling-is-all-about/", "Easy"},
			}},
		},
	}
}

func systemDesignSheet() sheetData {
	return sheetData{
		Name:        "System Design Sheet",
		Slug:        "system-design-sheet",
		Description: "Master HLD from Basics to Advanced",
		Topics: []topicData{
			{Name: "Basics", Problems: []problemData{
				{"What is System Design?", "https://www.geeksforgeeks.org/what-is-system-design-learn-system-design/", "Easy"},
				{"Horizontal vs Vertical Scaling", "https://www.geeksforgeeks.org/system-design-horizontal-and-vertical-scaling/", "Easy"},
				{"What is Capacity Estimation?", "https://www.geeksforgeeks.org/capacity-estimation/", "Easy"},
				{"What is HTTP?", "https://www.geeksforgeeks.org/http-full-form/", "Easy"},
				{"What is the Internet TCP/IP Stack?", "https://www.geeksforgeeks.org/tcp-ip-model/", "Medium"},
				{"What happens when you enter Google.com?", "https://www.geeksforgeeks.org/what-happens-when-you-type-a-url-in-your-browser/", "Medium"},
				{"What are Relational Databases?", "https://www.geeksforgeeks.org/relational-model-in-dbms/", "Easy"},
				{"What are Database Indexes?", "https://www.geeksforgeeks.org/indexing-in-databases-set-1/", "Medium"},
				{"What are NoSQL Databases?", "https://www.geeksforgeeks.org/introduction-to-nosql/", "Medium"},
				{"What is a Cache?", "https://www.geeksforgeeks.org/caching-system-design-concept-for-beginners/", "Easy"},
				{"What is Thrashing?", "https://www.geeksforgeeks.org/techniques-to-handle-thrashing/", "Medium"},
				{"What are Threads?", "https://www.geeksforgeeks.org/thread-in-operating-system/", "Easy"},
			}},
			{Name: "Load Balancing", Problems: []problemData{
				{"What is Load Balancing?", "https://www.geeksforgeeks.org/load-balancing-in-system-design/", "Medium"},
				{"What is Consistent Hashing?", "https://www.geeksforgeeks.org/consistent-hashing/", "Hard"},
				{"What is Sharding?", "https://www.geeksforgeeks.org/database-sharding-a-system-design-concept/", "Medium"},
			}},
			{Name: "DataStores", Problems: []problemData{
				{"What are Bloom Filters?", "https://www.geeksforgeeks.org/bloom-filters-introduction-and-python-implementation/", "Hard"},
				{"What is Data Replication?", "https://www.geeksforgeeks.org/data-replication-in-dbms/", "Medium"},
				{"How are NoSQL Databases Optimized?", "https://www.geeksforgeeks.org/introduction-to-nosql/", "Medium"},
				{"What are Location-based Databases?", "https://www.geeksforgeeks.org/spatial-databases/", "Medium"},
				{"Database Migrations", "https://www.geeksforgeeks.org/database-migration/", "Medium"},
			}},
			{Name: "Consistency vs Availability", Problems: []problemData{
				{"What is Data Consistency?", "https://www.geeksforgeeks.org/consistency-in-database/", "Medium"},
				{"Data Consistency Levels", "https://www.geeksforgeeks.org/eventual-vs-strong-consistency-in-distributed-databases/", "Hard"},
				{"Transaction Isolation Levels", "https://www.geeksforgeeks.org/transaction-isolation-levels-dbms/", "Hard"},
			}},
			{Name: "Message Queues", Problems: []problemData{
				{"What is a Message Queue?", "https://www.geeksforgeeks.org/message-queues-system-design/", "Medium"},
				{"Publisher-Subscriber Model", "https://www.geeksforgeeks.org/publish-subscribe-model/", "Medium"},
				{"Event-Driven Systems", "https://www.geeksforgeeks.org/event-driven-architecture-system-design/", "Medium"},
				{"Database as a Message Queue", "https://www.geeksforgeeks.org/message-queues-system-design/", "Medium"},
			}},
			{Name: "Caching", Problems: []problemData{
				{"What is Distributed Caching?", "https://www.geeksforgeeks.org/distributed-caching/", "Medium"},
				{"Content Delivery Networks (CDN)", "https://www.geeksforgeeks.org/what-is-a-content-delivery-network-and-how-does-it-work/", "Medium"},
				{"Write Policies", "https://www.geeksforgeeks.org/write-through-and-write-back-in-cache/", "Medium"},
				{"Replacement Policies", "https://www.geeksforgeeks.org/cache-replacement-policies/", "Medium"},
			}},
			{Name: "Microservices", Problems: []problemData{
				{"Microservices vs Monoliths", "https://www.geeksforgeeks.org/monolithic-vs-microservices-architecture/", "Medium"},
				{"How Monoliths are Migrated", "https://www.geeksforgeeks.org/monolithic-vs-microservices-architecture/", "Medium"},
			}},
			{Name: "System Design Practice Problems", Problems: []problemData{
				{"Design URL Shortener (TinyURL)", "https://www.geeksforgeeks.org/system-design-url-shortening-service/", "Medium"},
				{"Design Instagram", "https://www.geeksforgeeks.org/design-instagram/", "Hard"},
				{"Design WhatsApp", "https://www.geeksforgeeks.org/design-whatsapp/", "Hard"},
				{"Design Twitter", "https://www.geeksforgeeks.org/design-twitter-a-system-design-interview-question/", "Hard"},
				{"Design Netflix", "https://www.geeksforgeeks.org/system-design-netflix-a-complete-architecture/", "Hard"},
				{"Design Uber", "https://www.geeksforgeeks.org/system-design-of-uber-app-uber-system-architecture/", "Hard"},
				{"Design Amazon", "https://www.geeksforgeeks.org/system-design-of-amazon/", "Hard"},
				{"Design Google Maps", "https://www.geeksforgeeks.org/system-design-of-google-maps/", "Hard"},
				{"Design Google Docs", "https://www.geeksforgeeks.org/design-google-docs/", "Hard"},
			}},
		},
	}
}
