import 'package:flutter/material.dart';
import 'package:roadtrips/screens/suggestions.dart';
import 'package:url_launcher/url_launcher.dart';

class DestinationDetailsScreen extends StatefulWidget {
  static (String, Widget Function(BuildContext)) route =
      ('/destination/details', (context) => const DestinationDetailsScreen());

  const DestinationDetailsScreen({super.key});

  @override
  State<DestinationDetailsScreen> createState() => _DestinationDetailsState();
}

class _DestinationDetailsState extends State<DestinationDetailsScreen> {
  bool showBottomSheet = false;

  @override
  Widget build(BuildContext context) {
    final destination =
        ModalRoute.of(context)!.settings.arguments as Destination;

    return Scaffold(
      appBar: AppBar(
        title: const Text('Destination Details'),
        actions: [
          IconButton(
            icon: const Icon(Icons.more_vert),
            onPressed: () {
              showModalBottomSheet(
                  context: context,
                  builder: (context) {
                    return Container(
                      height: 80,
                      decoration: BoxDecoration(
                        borderRadius: const BorderRadius.only(
                          topLeft: Radius.circular(25),
                          topRight: Radius.circular(25),
                        ),
                        gradient: LinearGradient(
                          colors: [Colors.grey[300]!, Colors.grey[200]!],
                          begin: Alignment.topCenter,
                          end: Alignment.bottomCenter,
                          stops: const [.35, 0],
                        ),
                      ),
                      child: Column(
                        children: [
                          Container(
                            margin: const EdgeInsets.only(top: 12.5),
                            height: 3,
                            width: 30,
                            decoration: BoxDecoration(
                                color: Colors.white,
                                borderRadius: BorderRadius.circular(5)),
                          ),
                          GestureDetector(
                            onTap: () {
                              var uri = Uri.parse(destination.wazeDeepLink);
                              launchUrl(uri,
                                  mode:
                                      LaunchMode.externalNonBrowserApplication);
                            },
                            child: Container(
                              margin: const EdgeInsets.only(top: 9.5),
                              child: Container(
                                margin: const EdgeInsets.only(
                                    top: 17.5, left: 17.5),
                                alignment: Alignment.centerLeft,
                                child: const Text(
                                  'Open in Waze',
                                  textScaleFactor: 1.15,
                                ),
                              ),
                            ),
                          ),
                        ],
                      ),
                    );
                  });
            },
          ),
        ],
      ),
      body: Column(
        mainAxisAlignment: MainAxisAlignment.spaceAround,
        children: [
          Text(
            destination.name,
            style: const TextStyle(
              fontSize: 24,
              fontWeight: FontWeight.bold,
            ),
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceAround,
            children: [
              Column(
                children: [
                  Text(
                    'Latitude',
                    style: TextStyle(
                      color: Colors.grey[700]!,
                    ),
                  ),
                  Text(
                    destination.latitude.toString(),
                    style: const TextStyle(fontWeight: FontWeight.w500),
                  ),
                ],
              ),
              Column(
                children: [
                  Text(
                    'Longitude',
                    style: TextStyle(
                      color: Colors.grey[700]!,
                    ),
                  ),
                  Text(
                    destination.longitude.toString(),
                    style: const TextStyle(fontWeight: FontWeight.w500),
                  ),
                ],
              )
            ],
          )
        ],
      ),
    );
  }
}
