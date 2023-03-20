#!/bin/bash
CONFIG_PATH="./$HOME/.defund/config/config.toml"

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  sed -i 's/max_num_inbound_peers =.*/max_num_inbound_peers = 150/g' $CONFIG_PATH
  sed -i 's/max_num_outbound_peers =.*/max_num_outbound_peers = 150/g' $CONFIG_PATH
  sed -i 's/max_packet_msg_payload_size =.*/max_packet_msg_payload_size = 10240/g' $CONFIG_PATH
  sed -i 's/send_rate =.*/send_rate = 20480000/g' $CONFIG_PATH
  sed -i 's/recv_rate =.*/recv_rate = 20480000/g' $CONFIG_PATH
  sed -i 's/timeout_prevote =.*/timeout_prevote = "130ms"/g' $CONFIG_PATH
  sed -i 's/timeout_precommit =.*/timeout_precommit = "130ms"/g' $CONFIG_PATH
  sed -i 's/timeout_commit =.*/timeout_commit = "130ms"/g' $CONFIG_PATH
  sed -i 's/skip_timeout_commit =.*/skip_timeout_commit = false/g' $CONFIG_PATH
elif [[ "$OSTYPE" == "darwin"* ]]; then
  sed -i '' 's/max_num_inbound_peers =.*/max_num_inbound_peers = 150/g' $CONFIG_PATH
  sed -i '' 's/max_num_outbound_peers =.*/max_num_outbound_peers = 150/g' $CONFIG_PATH
  sed -i '' 's/max_packet_msg_payload_size =.*/max_packet_msg_payload_size = 10240/g' $CONFIG_PATH
  sed -i '' 's/send_rate =.*/send_rate = 20480000/g' $CONFIG_PATH
  sed -i '' 's/recv_rate =.*/recv_rate = 20480000/g' $CONFIG_PATH
  sed -i '' 's/timeout_prevote =.*/timeout_prevote = "130ms"/g' $CONFIG_PATH
  sed -i '' 's/timeout_precommit =.*/timeout_precommit = "130ms"/g' $CONFIG_PATH
  sed -i '' 's/timeout_commit =.*/timeout_commit = "130ms"/g' $CONFIG_PATH
  sed -i '' 's/skip_timeout_commit =.*/skip_timeout_commit = false/g' $CONFIG_PATH
else
  printf "Platform not supported, please ensure that the following values are set in your config.toml:\n"
  printf "###           P2P Configuration Options             ###\n"
  printf "\t max_num_inbound_peers = 150\n"
  printf "\t max_num_outbound_peers = 150\n"
  printf "\t max_packet_msg_payload_size = 10240\n"
  printf "\t send_rate = 20480000\n"
  printf "\t recv_rate = 20480000\n"
  printf "###         Consensus Configuration Options         ###\n"
  printf "\t timeout_prevote = \"130ms\"\n"
  printf "\t timeout_precommit = \"130ms\"\n"
  printf "\t timeout_commit = \"130ms\"\n"
  printf "\t skip_timeout_commit = false\n"
  exit 1
fi