case $1 in
  "RESET")
    rm -rf db
    mkdir db
    go install
    ;;
  "STOP")    
    pkill -f "dex"
    ;;
  "MC")
    nchainz create ETH 100000000 5 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz create USD 100000000 2 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    ;;
  "MAKEMARKET")
    nchainz order 1 ETH 990 USD 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 1100 USD 1 ETH 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 1 ETH 990 USD 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 1005 USD 1 ETH 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 1 ETH 850 USD 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 1042 USD 1 ETH 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 1 ETH 934 USD 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 1023 USD 1 ETH 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 1 ETH 923 USD 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 1092 USD 1 ETH 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 1 ETH 929 USD 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 1105 USD 1 ETH 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    ;;
  "TAKEMARKET")
    nchainz order 1 ETH 1100 USD 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 990 USD 1 ETH 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 1 ETH 1090 USD 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 850 USD 1 ETH 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 1 ETH 1010 USD 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 900 USD 1 ETH 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 1 ETH 1039 USD 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 905 USD 1 ETH 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 1 ETH 1017 USD 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 957 USD 1 ETH 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 1 ETH 1063 USD 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 963 USD 1 ETH 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 1 ETH 1033 USD 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    nchainz order 910 USD 1 ETH 1CufpgmhVmV7fujYHqFCqUdJe5vwhcc96K
    ;;
  "MAINNET")
    printf "35.172.27.110:5000\n35.174.80.163:5000\n52.87.50.12:5000" > seeds.txt
    ;;
  "LOCALNET")
    printf "localhost:4000\nlocalhost:4001\nlocalhost:4002" > seeds.txt
    ;;
  *)
    echo "Command not found: try again"
    ;;
esac

