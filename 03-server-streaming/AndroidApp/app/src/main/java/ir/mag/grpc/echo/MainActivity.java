package ir.mag.grpc.echo;

import android.media.AudioAttributes;
import android.media.AudioFormat;
import android.media.AudioManager;
import android.media.AudioTrack;
import android.media.MediaPlayer;
import android.os.Build;
import android.os.Bundle;
import android.util.Log;

import androidx.appcompat.app.AppCompatActivity;

import com.google.android.material.button.MaterialButton;

import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.util.Iterator;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class MainActivity extends AppCompatActivity {

	private static final String TAG = "MainActivity";

//	private AudioTrack myTrack;
//	private void initTrack() {
//		if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.LOLLIPOP) {
//			Log.d(TAG, "onCreate: going to create audio attr");
//			AudioAttributes audioAttributes = new AudioAttributes.Builder()
//					.setUsage(AudioAttributes.USAGE_MEDIA)
//					.setContentType(AudioAttributes.CONTENT_TYPE_MUSIC)
//					.build();
//
//			if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.P) {
//				Log.d(TAG, "onCreate: to create audio track");
//				myTrack = new AudioTrack(
//						audioAttributes,
//						new AudioFormat.Builder().build(),
//						200000,
//						AudioTrack.MODE_STREAM,
//						0);
//			}
//
//		}
//	}
//

	private MediaPlayer mediaPlayer = new MediaPlayer();
	private File tempMp3;

	@Override
	protected void onStart() {
		super.onStart();

//		while (myTrack == null || myTrack.getState() == 0) {
//			initTrack();
//			if (myTrack != null) {
//				Log.d(TAG, "onStart: state:" + myTrack.getState());
//			}
//		}

//		mediaPlayer.setAudioStreamType(AudioManager.STREAM_MUSIC);
//		if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.LOLLIPOP) {
//			mediaPlayer.setAudioAttributes(
//					new AudioAttributes.Builder()
//							.setContentType(AudioAttributes.CONTENT_TYPE_MUSIC)
//							.setUsage(AudioAttributes.USAGE_MEDIA)
//							.build()
//			);
//		}

		try {
			tempMp3 = File.createTempFile("kurchina", "mp3", getCacheDir());
			tempMp3.deleteOnExit();
		} catch (IOException e) {
			e.printStackTrace();
		}

	}

	@Override
	protected void onCreate(Bundle savedInstanceState) {
		super.onCreate(savedInstanceState);
		setContentView(R.layout.activity_main);

		String host = "192.168.1.46";
		int port = 8888;
		ManagedChannel channel = ManagedChannelBuilder.forAddress(host, port).usePlaintext().build();

		MaterialButton connectButton = findViewById(R.id.connect);
		connectButton.setOnClickListener(view -> {
			Log.d(TAG, "onCreate: clicked");
			Log.d(TAG, "globalCounter: in the global counter");
			CounterParam request = CounterParam.newBuilder().build();
			Iterator<GlobalCounterResponse> features;
			Log.d(TAG, "globalCounter: set features");
			features = EchoGrpc.newBlockingStub(channel).globalCounter(request);

			Log.d(TAG, "globalCounter: starting the while loop");
			while (features.hasNext()) {
				Log.d(TAG, "globalCounter: waiting");
				GlobalCounterResponse feature = features.next();
				Log.d(TAG, "globalCounter: " + feature.getCounter());
			}
		});

		MaterialButton playButton = findViewById(R.id.play);
		playButton.setOnClickListener(view -> {
			Log.d(TAG, "onCreate: click on play button");
			Log.d(TAG, "globalRadio: in the global radio");
			RadioParams request = RadioParams.newBuilder().build();
			Iterator<GlobalRadioResponse> features;
			Log.d(TAG, "globalRadio: set features");
			features = EchoGrpc.newBlockingStub(channel).globalRadio(request);

			Log.d(TAG, "globalRadio: starting the while loop");
			FileOutputStream fos = null;
			try {
				fos = new FileOutputStream(tempMp3);
			} catch (FileNotFoundException e) {
				e.printStackTrace();
			}
			while (features.hasNext()) {
				Log.d(TAG, "globalRadio: waiting");
				GlobalRadioResponse feature = features.next();
				byte[] byteArray = feature.getMusicBytes().toByteArray();
				try {
					fos.write(byteArray);
				} catch (IOException e) {
					e.printStackTrace();
				}
				Log.d(TAG, "globalRadio: " + feature.getMusicBytes());

				Log.d(TAG, "onCreate: track index is=" + feature.getChunkIndex());
				if (feature.getChunkIndex() == 0) {
					// create temp file that will hold byte array
//					try {
//
//						// resetting mediaplayer instance to evade problems
//						mediaPlayer.reset();
//
//						// In case you run into issues with threading consider new instance like:
//						// MediaPlayer mediaPlayer = new MediaPlayer();
//
//						// Tried passing path directly, but kept getting
//						// "Prepare failed.: status=0x1"
//						// so using file descriptor instead
//						FileInputStream fis = new FileInputStream(tempMp3);
//						mediaPlayer.setDataSource(fis.getFD());
//
//						mediaPlayer.prepare();
//						mediaPlayer.start();
//					} catch (IOException e) {
//						e.printStackTrace();
//					}
				}
			}

			try {
				fos.close();
			} catch (IOException e) {
				e.printStackTrace();
			}

			try {
				// resetting mediaplayer instance to evade problems
				mediaPlayer.reset();

				// In case you run into issues with threading consider new instance like:
				// MediaPlayer mediaPlayer = new MediaPlayer();

				// Tried passing path directly, but kept getting
				// "Prepare failed.: status=0x1"
				// so using file descriptor instead
				FileInputStream fis = new FileInputStream(tempMp3);
				mediaPlayer.setDataSource(fis.getFD());

				mediaPlayer.prepare();

				mediaPlayer.start();
			} catch (IOException e) {
				e.printStackTrace();
			}
		});

		Log.d(TAG, "onCreate: build sdk is " + Build.VERSION.SDK_INT);
	}

}